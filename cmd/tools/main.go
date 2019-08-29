package main

import (
	"fmt"
	"reflect"

	"knative.dev/serving/pkg/apis/serving/v1beta1"
)

func main() {
	item := v1beta1.Service{}
	imports := map[string]string{}
	t := reflect.TypeOf(item)
	tm := InterfacesToWrite(t, "knative.dev/serving/pkg/apis/serving/v1beta1", imports)
	interfaces := []string{}
	for tt, _ := range tm {
		interfaces = append(interfaces, WriteInterface(tt, imports, tm))
	}
	fmt.Printf(WriteImports(imports))
	for _, inter := range interfaces {
		fmt.Printf(inter)
	}
}
