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

package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	flag "github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/util/sets"
	"knative.dev/serving/pkg/apis/serving/v1alpha1"
	"knative.dev/serving/pkg/apis/serving/v1beta1"
)

var license = `// Copyright © 2019 The Knative Authors
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
`

func main() {
	api := flag.String("api", "", "API to generate for")
	out := flag.StringP("out", "o", "", "File to write")
	inter := flag.Bool("interface", false, "Generate the interface")
	var interPkg, myPkg string
	flag.StringVar(&interPkg, "interface-package", "", "Package for the interface")
	flag.StringVar(&myPkg, "my-package", "", "Package for the file to be written")

	flag.Parse()

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
		PkgPath:      "knative.dev/serving/pkg/apis/serving/v1beta1",
		InterfacePkg: interPkg,
		MyPkg:        myPkg,
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
	betaCtx := op.NewAbstractionContext(nil)
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
		InterfacePkg: interPkg,
		MyPkg:        myPkg,
		PkgPath:      "knative.dev/serving/pkg/apis/serving/v1alpha1",
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
	alphaCtx := alphaOp.NewAbstractionContext(nil)
	contexts := map[string]*AbstractionContext{
		"v1alpha1": alphaCtx,
		"v1beta1":  betaCtx,
	}
	f, err := os.OpenFile(*out, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		dir, _ := os.Getwd()
		fmt.Printf("Current dir %s\n", dir)
		fmt.Printf("Trying to create %s\n", *out)
		panic(err)
	}
	fmt.Fprintf(f, "%s\n// Code generated by go run ./cmd/tools %s. DO NOT EDIT.\n", license, strings.Join(os.Args[1:], " "))
	ctx := contexts[*api]

	if *inter {
		ctx.WriteInterfaceFile(f)
	} else {
		ctx.WriteImplementationFile(f)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
