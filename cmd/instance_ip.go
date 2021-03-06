package cmd

import (
	"fmt"
	"os"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var instancePublicIPCmd = &cobra.Command{
	Use:     "public-ip",
	Short:   "Show instance's public IP",
	Aliases: []string{"ip", "publicip"},
	Long: `Show the specified instance's public IP by part of the instance's ID or name.
If you wish to use a custom format, the available fields are:

	* ID
	* Hostname
	* PublicIP

Example: civo instance public-ip ID/NAME`,
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

		if outputFormat == "human" {
			fmt.Printf("The instance %s (%s) has the public IP %s\n", aurora.Green(instance.Hostname), instance.ID, aurora.Green(instance.PublicIP))
		} else {
			ow := utility.NewOutputWriter()
			ow.StartLine()
			ow.AppendData("ID", instance.ID)
			ow.AppendData("Hostname", instance.Hostname)
			ow.AppendDataWithLabel("PublicIP", instance.PublicIP, "Public ID")
			if outputFormat == "json" {
				ow.WriteSingleObjectJSON()
			} else {
				ow.WriteCustomOutput(outputFields)
			}
		}
	},
}
