package cmd

import (
	"fmt"
	"os"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var instanceRebootCmd = &cobra.Command{
	Use:     "reboot",
	Aliases: []string{"hard-reboot"},
	Short:   "Hard reboot an instance",
	Long: `Pull the power and restart the specified instance by part of the ID or name.
If you wish to use a custom format, the available fields are:

	* ID
	* Hostname

Example: civo instance reboot ID/NAME`,
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

		_, err = client.RebootInstance(instance.ID)
		if err != nil {
			fmt.Printf("Rebooting instance: %s\n", aurora.Red(err))
			os.Exit(1)
		}

		if outputFormat == "human" {
			fmt.Printf("The instance %s (%s) is being rebooted\n", aurora.Green(instance.Hostname), instance.ID)
		} else {
			ow := utility.NewOutputWriter()
			ow.StartLine()
			ow.AppendData("ID", instance.ID)
			ow.AppendData("Hostname", instance.Hostname)
			if outputFormat == "json" {
				ow.WriteSingleObjectJSON()
			} else {
				ow.WriteCustomOutput(outputFields)
			}
		}
	},
}
