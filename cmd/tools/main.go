package main

import (
	"fmt"
	"os"
	"reflect"

	"k8s.io/apimachinery/pkg/util/sets"

	"knative.dev/serving/pkg/apis/serving/v1beta1"
)

func main() {

	nontemplateMetadataFields := sets.NewString("GenerateName", "Namespace", "SelfLink", "UID", "ResourceVersion", "Generation", "CreationTimestamp",
		"DeletionTimestamp", "DeletionGracePeriodSeconds", "OwnerReferences", "Initializers", "Finalizers", "ClusterName")
	op := GenerationOptions{
		OmitFields: map[string]sets.String{
			"RevisionSpec": sets.NewString("InitContainers", "RestartPolicy", "TerminationGracePeriodSeconds", "ActiveDeadlineSeconds", "DNSPolicy",
				"NodeSelector", "DeprecatedServiceAccount", "AutomountServiceAccountToken", "NodeName", "HostNetwork", "HostPID",
				"HostIPC", "ShareProcessNamespace", "SecurityContext", "ImagePullSecrets", "Hostname", "Subdomain", "Affinity",
				"SchedulerName", "Tolerations", "HostAliases", "PriorityClassName", "Priority", "DNSConfig", "ReadinessGates",
				"RuntimeClassName"),
			"RevisionTemplateSpec": nontemplateMetadataFields,
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
	ctx := op.NewAbstractionContext()
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
	fmt.Println("package generic\n")
	fmt.Printf(ctx.WriteImports())
	for _, d := range declarations {
		d.WriteDeclaration(os.Stdout)
	}
}
