// Copyright © 2019 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"knative.dev/serving/pkg/apis/serving"

	"knative.dev/client/pkg/printers"
	clientserving "knative.dev/client/pkg/serving"
	serving_kn_v1alpha1 "knative.dev/client/pkg/serving/v1alpha1"
	"knative.dev/client/pkg/wait"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"knative.dev/client/pkg/kn/commands"
	"knative.dev/serving/pkg/apis/serving/v1alpha1"
)

// Command for printing out a description of a service, meant to be consumed by humans
// It will show information about the service itself, but also a summary
// about the associated revisions.

// Whether to print extended information
var printDetails bool

// Matching image digest
var imageDigestRegexp = regexp.MustCompile(`(?i)sha256:([0-9a-f]{64})`)

// View object for collecting revision related information in the context
// of a Service. These are plain data types which can be directly used
// for printing out
type revisionDesc struct {
	revision *v1alpha1.Revision

	// traffic stuff
	percent       int
	tag           string
	latestTraffic *bool

	configurationGeneration int

	// status info
	latestCreated bool
	latestReady   bool
}

// [REMOVE COMMENT WHEN MOVING TO 0.7.0]
// For transition to v1beta1 this command uses the migration approach as described
// in https://docs.google.com/presentation/d/1mOhnhy8kA4-K9Necct-NeIwysxze_FUule-8u5ZHmwA/edit#slide=id.p
// With serving 0.6.0 we are at step #1
// I.e we first look at new fields of the v1alpha1 API before falling back to the original ones.
// As this command does not do any writes/updates, it's just a matter of fallbacks.
// [/REMOVE COMMENT WHEN MOVING TO 0.7.0]

// NewServiceDescribeCommand returns a new command for describing a service.
func NewServiceDescribeCommand(p *commands.KnParams) *cobra.Command {

	// For machine readable output
	machineReadablePrintFlags := genericclioptions.NewPrintFlags("")

	command := &cobra.Command{
		Use:   "describe NAME",
		Short: "Show details for a given service",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("no service name provided")
			}
			if len(args) > 1 {
				return errors.New("more than one service name provided")
			}
			serviceName := args[0]

			namespace, err := p.GetNamespace(cmd)
			if err != nil {
				return err
			}

			client, err := p.NewClient(namespace)
			if err != nil {
				return err
			}

			service, err := client.GetService(serviceName)
			if err != nil {
				return err
			}

			// Print out machine readable output if requested
			if machineReadablePrintFlags.OutputFlagSpecified() {
				printer, err := machineReadablePrintFlags.ToPrinter()
				if err != nil {
					return err
				}
				return printer.PrintObj(service, cmd.OutOrStdout())
			}

			printDetails, err = cmd.Flags().GetBool("verbose")
			if err != nil {
				return err
			}

			revisionDescs, err := getRevisionDescriptions(client, service, printDetails)
			if err != nil {
				return err
			}

			return describe(cmd.OutOrStdout(), service, revisionDescs, printDetails)
		},
	}
	flags := command.Flags()
	commands.AddNamespaceFlags(flags, false)
	flags.BoolP("verbose", "v", false, "More output.")
	machineReadablePrintFlags.AddFlags(command)
	return command
}

// Main action describing the service
func describe(w io.Writer, service *v1alpha1.Service, revisions []*revisionDesc, printDetails bool) error {
	dw := printers.NewPrefixWriter(w)

	// Service info
	writeService(dw, service)
	dw.WriteLine()
	if err := dw.Flush(); err != nil {
		return err
	}

	// Revisions summary info
	writeRevisions(dw, revisions, printDetails)
	dw.WriteLine()
	if err := dw.Flush(); err != nil {
		return err
	}

	// Condition info
	commands.WriteConditions(dw, service.Status.Conditions, printDetails)
	if err := dw.Flush(); err != nil {
		return err
	}

	return nil
}

// Write out main service information. Use colors for major items.
func writeService(dw printers.PrefixWriter, service *v1alpha1.Service) {
	commands.WriteMetadata(dw, &service.ObjectMeta, printDetails)
	dw.WriteAttribute("URL", extractURL(service))
	if service.Status.Address != nil {
		url := service.Status.Address.GetURL()
		dw.WriteAttribute("Address", url.String())
	}
}

// Write out revisions associated with this service. By default only active
// target revisions are printed, but with --all also inactive revisions
// created by this services are shown
func writeRevisions(dw printers.PrefixWriter, revisions []*revisionDesc, printDetails bool) {
	revSection := dw.WriteAttribute("Revisions", "")
	dw.Flush()
	for _, revisionDesc := range revisions {
		ready := wait.ReadyCondition(revisionDesc.revision.Status.Conditions)
		section := revSection.WriteColsLn(formatBullet(revisionDesc.percent, ready.Status), revisionHeader(revisionDesc))
		if ready.Status == v1.ConditionFalse {
			section.WriteAttribute("Error", ready.Reason)
		}
		section.WriteAttribute("Image", getImageDesc(revisionDesc))
		if printDetails {
			port := clientserving.Port(&revisionDesc.revision.Spec)
			if port != nil {
				section.WriteAttribute("Port", strconv.FormatInt(int64(*port), 10))
			}
			env := stringifyEnv(revisionDesc.revision)
			if env != nil {
				writeSliceDesc(section, env, l("Env"), "")
			}

			// Scale spec if given
			scale, _ := clientserving.ScalingInfo(&revisionDesc.revision.ObjectMeta, &revisionDesc.revision.Spec)
			if scale != nil && (scale.Max != nil || scale.Min != nil) {
				section.WriteAttribute("Scale", formatScale(scale.Min, scale.Max))
			}

			// Concurrency specs if given
			writeConcurrencyOptions(section, revisionDesc.revision)

			// Resources if given
			writeResources(section, revisionDesc.revision)
		}
	}
}

func writeConcurrencyOptions(dw printers.PrefixWriter, revision *v1alpha1.Revision) {
	target := clientserving.ConcurrencyTarget(&revision.ObjectMeta)
	limit := revision.Spec.ContainerConcurrency
	if target != nil || limit != 0 {
		section := dw.WriteAttribute("Concurrency", "")
		if limit != 0 {
			section.WriteAttribute("Limit", strconv.FormatInt(int64(limit), 10))
		}
		if target != nil {
			section.WriteAttribute("Target", strconv.Itoa(*target))
		}
	}
}

// ======================================================================================
// Helper functions

func stringifyEnv(revision *v1alpha1.Revision) []string {
	container, err := clientserving.ContainerOfRevisionSpec(&revision.Spec)
	if err != nil {
		return nil
	}

	envVars := make([]string, 0, len(container.Env))
	for _, env := range container.Env {
		var value string
		if env.ValueFrom != nil {
			value = "[ref]"
		} else {
			value = env.Value
		}
		envVars = append(envVars, fmt.Sprintf("%s=%s", env.Name, value))
	}
	return envVars
}

// Format label (extracted so that color could be added more easily to all labels)
func l(label string) string {
	return label + ":"
}

// Format scale in the format "min ... max" with max = ∞ if not set
func formatScale(minScale *int, maxScale *int) string {
	ret := "0"
	if minScale != nil {
		ret = strconv.Itoa(*minScale)
	}

	ret += " ... "

	if maxScale != nil {
		ret += strconv.Itoa(*maxScale)
	} else {
		ret += "∞"
	}
	return ret
}

// Format the revision name along with its generation. Use colors if enabled.
func revisionHeader(desc *revisionDesc) string {
	header := desc.revision.Name
	if desc.latestTraffic != nil && *desc.latestTraffic {
		header = fmt.Sprintf("@latest (%s)", desc.revision.Name)
	} else if desc.latestReady {
		header = desc.revision.Name + " (current @latest)"
	} else if desc.latestCreated {
		header = desc.revision.Name + " (latest created)"
	}
	if desc.tag != "" {
		header = fmt.Sprintf("%s #%s", header, desc.tag)
	}
	return header + " " +
		"[" + strconv.Itoa(desc.configurationGeneration) + "]" +
		" " +
		"(" + commands.Age(desc.revision.CreationTimestamp.Time) + ")"
}

// Return either image name with tag or together with its resolved digest
func getImageDesc(desc *revisionDesc) string {
	c, err := clientserving.ContainerOfRevisionSpec(&desc.revision.Spec)
	if err != nil {
		return "Unknown"
	}
	image := c.Image
	// Check if the user image is likely a more user-friendly description
	pinnedDesc := "at"
	userImage := clientserving.UserImage(&desc.revision.ObjectMeta)
	imageDigest := desc.revision.Status.ImageDigest
	if userImage != "" && imageDigest != "" {
		parts := strings.Split(image, "@")
		// Check if the user image refers to the same thing.
		if strings.HasPrefix(userImage, parts[0]) {
			pinnedDesc = "pinned to"
			image = userImage
		}
	}
	if imageDigest != "" {
		return fmt.Sprintf("%s (%s %s)", image, pinnedDesc, shortenDigest(imageDigest))
	}
	return image
}

// Extract pure sha sum and shorten to 8 digits,
// as the digest should to be user consumable. Use the resource via `kn service get`
// to get to the full sha
func shortenDigest(digest string) string {
	match := imageDigestRegexp.FindStringSubmatch(digest)
	if len(match) > 1 {
		return string(match[1][:6])
	}
	return digest
}

// Writer a slice compact (printDetails == false) in one line, or over multiple line
// with key-value line-by-line (printDetails == true)
func writeSliceDesc(dw printers.PrefixWriter, s []string, label string, labelPrefix string) {

	if len(s) == 0 {
		return
	}

	if printDetails {
		l := labelPrefix + label
		for _, value := range s {
			dw.WriteColsLn(l, value)
			l = labelPrefix
		}
		return
	}

	joined := strings.Join(s, ", ")
	if len(joined) > commands.TruncateAt {
		joined = joined[:commands.TruncateAt-4] + " ..."
	}
	dw.WriteAttribute(labelPrefix+label, joined)
}

func writeResources(dw printers.PrefixWriter, r *v1alpha1.Revision) {
	c, err := clientserving.ContainerOfRevisionSpec(&r.Spec)
	if err != nil {
		return
	}
	requests := c.Resources.Requests
	limits := c.Resources.Limits
	writeResourcesHelper(dw, "Memory", requests.Memory(), limits.Memory())
	writeResourcesHelper(dw, "CPU", requests.Cpu(), limits.Cpu())
}

// Write request ... limits or only one of them
func writeResourcesHelper(dw printers.PrefixWriter, label string, request *resource.Quantity, limit *resource.Quantity) {
	value := ""
	if !request.IsZero() && !limit.IsZero() {
		value = request.String() + " ... " + limit.String()
	} else if !request.IsZero() {
		value = request.String()
	} else if !limit.IsZero() {
		value = limit.String()
	}

	if value == "" {
		return
	}

	dw.WriteAttribute(label, value)
}

// Format target percentage that it fits in the revision table
func formatBullet(percentage int, status corev1.ConditionStatus) string {
	symbol := "+"
	switch status {
	case v1.ConditionTrue:
		if percentage > 0 {
			symbol = "%"
		}
	case v1.ConditionFalse:
		symbol = "!"
	default:
		symbol = "?"
	}
	if percentage == 0 {
		return fmt.Sprintf("   %s", symbol)
	}
	return fmt.Sprintf("%3d%s", percentage, symbol)
}

// Call the backend to query revisions for the given service and build up
// the view objects used for output
func getRevisionDescriptions(client serving_kn_v1alpha1.KnServingClient, service *v1alpha1.Service, withDetails bool) ([]*revisionDesc, error) {
	revisionsSeen := sets.NewString()
	revisionDescs := []*revisionDesc{}

	trafficTargets := service.Status.Traffic
	var err error
	for _, target := range trafficTargets {
		revision, err := extractRevisionFromTarget(client, target)
		if err != nil {
			return nil, fmt.Errorf("cannot extract revision from service %s: %v", service.Name, err)
		}
		revisionsSeen.Insert(revision.Name)
		desc, err := newRevisionDesc(revision, &target, service)
		if err != nil {
			return nil, err
		}
		revisionDescs = append(revisionDescs, desc)
	}
	if revisionDescs, err = completeWithLatestRevisions(client, service, revisionsSeen, revisionDescs); err != nil {
		return nil, err
	}
	if withDetails {
		if revisionDescs, err = completeWithUntargetedRevisions(client, service, revisionsSeen, revisionDescs); err != nil {
			return nil, err
		}
	}
	return orderByConfigurationGeneration(revisionDescs), nil
}

// Order the list of revisions so that the newest revisions are at the top
func orderByConfigurationGeneration(descs []*revisionDesc) []*revisionDesc {
	sort.SliceStable(descs, func(i, j int) bool {
		return descs[i].configurationGeneration > descs[j].configurationGeneration
	})
	return descs
}

func completeWithLatestRevisions(client serving_kn_v1alpha1.KnServingClient, service *v1alpha1.Service, revisionsSeen sets.String, descs []*revisionDesc) ([]*revisionDesc, error) {
	for _, revisionName := range []string{service.Status.LatestCreatedRevisionName, service.Status.LatestReadyRevisionName} {
		if revisionsSeen.Has(revisionName) {
			continue
		}
		revisionsSeen.Insert(revisionName)
		rev, err := client.GetRevision(revisionName)
		if err != nil {
			return nil, err
		}
		newDesc, err := newRevisionDesc(rev, nil, service)
		if err != nil {
			return nil, err
		}
		descs = append(descs, newDesc)
	}
	return descs, nil
}

func completeWithUntargetedRevisions(client serving_kn_v1alpha1.KnServingClient, service *v1alpha1.Service, revisionsSeen sets.String, descs []*revisionDesc) ([]*revisionDesc, error) {
	revisions, err := client.ListRevisions(serving_kn_v1alpha1.WithService(service.Name))
	if err != nil {
		return nil, err
	}
	for _, revision := range revisions.Items {
		if revisionsSeen.Has(revision.Name) {
			continue
		}
		revisionsSeen.Insert(revision.Name)
		newDesc, err := newRevisionDesc(&revision, nil, service)
		if err != nil {
			return nil, err
		}
		descs = append(descs, newDesc)

	}
	return descs, nil
}

func newRevisionDesc(revision *v1alpha1.Revision, target *v1alpha1.TrafficTarget, service *v1alpha1.Service) (*revisionDesc, error) {
	generation, err := strconv.ParseInt(revision.Labels[serving.ConfigurationGenerationLabelKey], 0, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot extract configuration generation for revision %s: %v", revision.Name, err)
	}
	revisionDesc := revisionDesc{
		revision:                revision,
		configurationGeneration: int(generation),
		latestCreated:           revision.Name == service.Status.LatestCreatedRevisionName,
		latestReady:             revision.Name == service.Status.LatestReadyRevisionName,
	}

	addTargetInfo(&revisionDesc, target)
	if err != nil {
		return nil, err
	}
	return &revisionDesc, nil
}

func addTargetInfo(desc *revisionDesc, target *v1alpha1.TrafficTarget) {
	if target != nil {
		desc.percent = target.Percent
		desc.latestTraffic = target.LatestRevision
		desc.tag = target.Tag
	}
}

func extractRevisionFromTarget(client serving_kn_v1alpha1.KnServingClient, target v1alpha1.TrafficTarget) (*v1alpha1.Revision, error) {
	var revisionName = target.RevisionName
	if revisionName == "" {
		configurationName := target.ConfigurationName
		if configurationName == "" {
			return nil, fmt.Errorf("neither RevisionName nor ConfigurationName set")
		}
		configuration, err := client.GetConfiguration(configurationName)
		if err != nil {
			return nil, err
		}
		revisionName = configuration.Status.LatestCreatedRevisionName
	}
	return client.GetRevision(revisionName)
}

func extractURL(service *v1alpha1.Service) string {
	status := service.Status
	if status.URL != nil {
		return status.URL.String()
	}
	return status.DeprecatedDomain
}
