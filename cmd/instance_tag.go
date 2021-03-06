package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var instanceTagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"tags"},
	Short:   "Change the instance's tags",
	Long: `Change the tags for an instance with partial ID/name provided.
If you wish to use a custom format, the available fields are:

* ID
* Hostname
* Tags

Example: civo instance tags ID/NAME tag1 tag2 tag3`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			fmt.Printf("Unable to create a Civo API Client: %s\n", aurora.Red(err))
			os.Exit(1)
		}

		instance, err := client.FindInstance(args[0])
		if err != nil {
			fmt.Printf("Finding instance: %s\n", aurora.Red(err))
			os.Exit(1)
		}

		tags := strings.Join(args[1:], " ")

		_, err = client.SetInstanceTags(instance, tags)
		if err != nil {
			fmt.Printf("Retagging instance: %s\n", aurora.Red(err))
			os.Exit(1)
		}

		if outputFormat == "human" {
			fmt.Printf("The instance %s (%s) has been tagged with '%s'\n", aurora.Green(instance.Hostname), instance.ID, aurora.Green(tags))
		} else {
			ow := utility.NewOutputWriter()
			ow.StartLine()
			ow.AppendData("ID", instance.ID)
			ow.AppendData("Hostname", instance.Hostname)
			ow.AppendDataWithLabel("ReverseDNS", instance.ReverseDNS, "Reverse DNS")
			ow.AppendData("Notes", instance.Notes)
			if outputFormat == "json" {
				ow.WriteSingleObjectJSON()
			} else {
				ow.WriteCustomOutput(outputFields)
			}
		}
	},
}
