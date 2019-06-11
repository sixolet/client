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

package plugin

import (
	"fmt"

	"github.com/knative/client/pkg/kn/commands"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NewPluginListCommand represents 'kn service get' command
func NewPluginListCommand(p *commands.KnParams) *cobra.Command {
	pluginListFlags := NewPluginListFlags()

	pluginListCommand := &cobra.Command{
		Use:   "get",
		Short: "List available plugins.",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := p.ClientFactory()
			if err != nil {
				return err
			}
			namespace, err := commands.GetNamespace(cmd)
			if err != nil {
				return err
			}
			plugin, err := client.Plugins(namespace).List(v1.ListOptions{})
			if err != nil {
				return err
			}
			if len(plugin.Items) == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "No resources found.\n")
				return nil
			}
			plugin.GetObjectKind().SetGroupVersionKind(schema.GroupVersionKind{
				Group:   "client.knative.dev",
				Version: "v1alpha1",
				Kind:    "Plugin"})

			printer, err := pluginListFlags.ToPrinter()
			if err != nil {
				return err
			}

			err = printer.PrintObj(plugin, cmd.OutOrStdout())
			if err != nil {
				return err
			}
			return nil
		},
	}
	commands.AddNamespaceFlags(pluginListCommand.Flags(), true)
	pluginListFlags.AddFlags(pluginListCommand)
	return pluginListCommand
}
