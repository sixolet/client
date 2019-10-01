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

package revision

import (
	"errors"
	"io"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"knative.dev/client/pkg/kn/commands"
	"knative.dev/client/pkg/printers"
	"knative.dev/serving/pkg/apis/serving/v1alpha1"
)

func NewRevisionDescribeCommand(p *commands.KnParams) *cobra.Command {

	// For machine readable output
	machineReadablePrintFlags := genericclioptions.NewPrintFlags("")

	command := &cobra.Command{
		Use:   "describe NAME",
		Short: "Describe revisions.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires the revision name.")
			}

			namespace, err := p.GetNamespace(cmd)
			if err != nil {
				return err
			}
			client, err := p.NewClient(namespace)
			if err != nil {
				return err
			}

			revision, err := client.GetRevision(args[0])
			if err != nil {
				return err
			}

			if machineReadablePrintFlags.OutputFlagSpecified() {
				printer, err := machineReadablePrintFlags.ToPrinter()
				if err != nil {
					return err
				}
				return printer.PrintObj(revision, cmd.OutOrStdout())
			}
			printDetails, err := cmd.Flags().GetBool("verbose")
			// Do the human-readable printing thing.
			return describe(cmd.OutOrStdout(), revision, printDetails)
		},
	}
	flags := command.Flags()
	commands.AddNamespaceFlags(flags, false)
	machineReadablePrintFlags.AddFlags(command)
	flags.BoolP("verbose", "v", false, "More output.")
	return command
}

func describe(w io.Writer, revision *v1alpha1.Revision, printDetails bool) error {
	dw := printers.NewPrefixWriter(w)

	commands.WriteConditions(dw, revision.Status.Conditions, printDetails)
	if err := dw.Flush(); err != nil {
		return err
	}
	return nil
}
