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
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or im
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	clientv1alpha1 "github.com/knative/client/pkg/apis/client/v1alpha1"
	"github.com/knative/client/pkg/kn/commands"
	hprinters "github.com/knative/client/pkg/printers"
	"github.com/spf13/cobra"
	metav1beta1 "k8s.io/apimachinery/pkg/apis/meta/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

// PluginListFlags composes common printer flag structs
// used in the List command.
type PluginListFlags struct {
	GenericPrintFlags  *genericclioptions.PrintFlags
	HumanReadableFlags *commands.HumanPrintFlags
}

// AllowedFormats is the list of formats in which data can be displayed
func (f *PluginListFlags) AllowedFormats() []string {
	formats := f.GenericPrintFlags.AllowedFormats()
	formats = append(formats, f.HumanReadableFlags.AllowedFormats()...)
	return formats
}

// ToPrinter attempts to find a composed set of PluginListFlags suitable for
// returning a printer based on current flag values.
func (f *PluginListFlags) ToPrinter() (hprinters.ResourcePrinter, error) {
	// if there are flags specified for generic printing
	if f.GenericPrintFlags.OutputFlagSpecified() {
		p, err := f.GenericPrintFlags.ToPrinter()
		if err != nil {
			return nil, err
		}
		return p, nil
	}
	// if no flags specified, use the table printing
	p, err := f.HumanReadableFlags.ToPrinter(PluginListHandlers)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// AddFlags receives a *cobra.Command reference and binds
// flags related to humanreadable and template printing.
func (f *PluginListFlags) AddFlags(cmd *cobra.Command) {
	f.GenericPrintFlags.AddFlags(cmd)
	f.HumanReadableFlags.AddFlags(cmd)
}

// NewListPrintFlags returns flags associated with humanreadable,
// template, and "name" printing, with default values set.
func NewPluginListFlags() *PluginListFlags {
	return &PluginListFlags{
		GenericPrintFlags:  genericclioptions.NewPrintFlags(""),
		HumanReadableFlags: commands.NewHumanPrintFlags(),
	}
}

// Human-readable columns for a Plugin
func PluginListHandlers(h hprinters.PrintHandler) {
	kPluginColumnDefinitions := []metav1beta1.TableColumnDefinition{
		{Name: "Name", Type: "string", Description: "Plugin name"},
		{Name: "Description", Type: "string", Description: "Plugin description"},
	}
	h.TableHandler(kPluginColumnDefinitions, printPlugin)
	h.TableHandler(kPluginColumnDefinitions, printPluginList)
}

// Private functions

func printPluginList(pluginList *clientv1alpha1.PluginList, options hprinters.PrintOptions) ([]metav1beta1.TableRow, error) {
	rows := make([]metav1beta1.TableRow, 0, len(pluginList.Items))
	for _, pl := range pluginList.Items {
		r, err := printPlugin(&pl, options)
		if err != nil {
			return nil, err
		}
		rows = append(rows, r...)
	}
	return rows, nil
}

// printKService populates the knative service table rows
func printPlugin(plugin *clientv1alpha1.Plugin, options hprinters.PrintOptions) ([]metav1beta1.TableRow, error) {
	name := plugin.Name
	description := plugin.Spec.Description

	row := metav1beta1.TableRow{
		Object: runtime.RawExtension{Object: plugin},
	}
	row.Cells = append(row.Cells,
		name,
		description)
	return []metav1beta1.TableRow{row}, nil
}
