package main

import (
	"fmt"
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
	imports := map[string]string{}
	tt := []reflect.Type{
		reflect.TypeOf(v1beta1.Service{}),
		reflect.TypeOf(v1beta1.Configuration{}),
		reflect.TypeOf(v1beta1.Route{}),
		reflect.TypeOf(v1beta1.Revision{}),
	}
	tm := InterfacesToWrite(tt, "knative.dev/serving/pkg/apis/serving/v1beta1", imports)
	interfaces := []string{}
	impls := []string{}
	for t, _ := range tm {
		switch t.Kind() {
		case reflect.Struct:
			interfaces = append(interfaces, WriteInterface(t, imports, tm, op))
			impls = append(impls, WriteImplementation(t, imports, tm, op))
		case reflect.Slice:
			interfaces = append(interfaces, WriteSliceInterface(t, imports, tm, op))
		}
	}
	fmt.Println("package generic\n")
	fmt.Printf(WriteImports(imports))
	for _, inter := range interfaces {
		fmt.Printf(inter)
	}
	for _, impl := range impls {
		fmt.Printf(impl)
	}
}
