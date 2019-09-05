package main

import (
	"fmt"
	"os"
	"reflect"

	"k8s.io/apimachinery/pkg/util/sets"
	"knative.dev/serving/pkg/apis/serving/v1alpha1"
	"knative.dev/serving/pkg/apis/serving/v1beta1"
)

func main() {

	nontemplateMetadataFields := sets.NewString("GenerateName", "Namespace", "SelfLink", "UID", "ResourceVersion", "Generation", "CreationTimestamp",
		"DeletionTimestamp", "DeletionGracePeriodSeconds", "OwnerReferences", "Initializers", "Finalizers", "ClusterName")
	revisionOmitFields := sets.NewString("InitContainers", "RestartPolicy", "TerminationGracePeriodSeconds", "ActiveDeadlineSeconds", "DNSPolicy",
		"NodeSelector", "DeprecatedServiceAccount", "AutomountServiceAccountToken", "NodeName", "HostNetwork", "HostPID",
		"HostIPC", "ShareProcessNamespace", "SecurityContext", "ImagePullSecrets", "Hostname", "Subdomain", "Affinity",
		"SchedulerName", "Tolerations", "HostAliases", "PriorityClassName", "Priority", "DNSConfig", "ReadinessGates",
		"RuntimeClassName")

	op := GenerationOptions{
		OmitFields: map[string]sets.String{
			"RevisionTemplateSpec": nontemplateMetadataFields,
			"RevisionSpec":         revisionOmitFields,
		},
		PkgPath: "knative.dev/serving/pkg/apis/serving/v1beta1",
		Roots: []reflect.Type{
			reflect.TypeOf(v1beta1.Service{}),
			reflect.TypeOf(v1beta1.Configuration{}),
			reflect.TypeOf(v1beta1.Route{}),
			reflect.TypeOf(v1beta1.Revision{}),
		},
		SliceKeys: map[string]sets.String{
			"TrafficTarget": sets.NewString("Tag", "RevisionName", "LatestRevision"),
		},
	}
	ctx := op.NewAbstractionContext(nil)
	declarations := []Declaration{}
	for t, _ := range ctx.Abstract {
		switch t.Kind() {
		case reflect.Struct:
			declarations = append(declarations, ctx.MakeInterface(t))
			declarations = append(declarations, ctx.MakeImplementation(t))
		case reflect.Slice:
			declarations = append(declarations, ctx.MakeSliceInterface(t))
			declarations = append(declarations, ctx.MakeSliceImpl(t))
		}
	}

	alphaOp := GenerationOptions{
		OmitFields: map[string]sets.String{
			"RevisionTemplateSpec": nontemplateMetadataFields,
			"RevisionSpec":         revisionOmitFields,
		},
		OmitImplementations: map[string]sets.String{
			"ServiceSpec":        sets.NewString("Template"),
			"ConfigurationSpec":  sets.NewString("Template"),
			"RevisionSpec":       sets.NewString("Containers"),
			"TrafficTargetSlice": sets.NewString("Upsert"),
			"ServiceStatus":      sets.NewString("Address"),
			"RouteStatus":        sets.NewString("Address"),
		},
		PkgPath: "knative.dev/serving/pkg/apis/serving/v1alpha1",
		Roots: []reflect.Type{
			reflect.TypeOf(v1alpha1.Service{}),
			reflect.TypeOf(v1alpha1.Configuration{}),
			reflect.TypeOf(v1alpha1.Route{}),
			reflect.TypeOf(v1alpha1.Revision{}),
		},
		SliceKeys: map[string]sets.String{
			"TrafficTarget": sets.NewString("Tag", "RevisionName", "LatestRevision"),
		},
	}
	alphaCtx := alphaOp.NewAbstractionContext(ctx.Imports)
	for t, _ := range alphaCtx.Abstract {
		switch t.Kind() {
		case reflect.Struct:
			declarations = append(declarations, alphaCtx.MakeImplementation(t))
		case reflect.Slice:

			declarations = append(declarations, alphaCtx.MakeSliceImpl(t))
		}
	}

	fmt.Println("package generic\n")
	fmt.Printf(ctx.WriteImports())
	for _, d := range declarations {
		d.WriteDeclaration(os.Stdout)
	}
}
