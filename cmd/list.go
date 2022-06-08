package cmd

import (
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.ExactArgs(0),
	Short: "A brief description of your command",
	Long:  `A longer description that spans`,
	Run:   list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(_ *cobra.Command, _ []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NAME", "IMAGE", "TAG", "PARAMETER", "COMMAND", "PATH"})
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetColumnSeparator("")
	table.SetHeaderLine(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(GetCommands())
	table.Render()
}

func GetCommands() [][]string {

	containers := config.Instance.Get(config.Commands)

	var data [][]string

	if !(containers == "{}") {
		for key := range containers.(map[string]interface{}) {
			data = append(data, []string{
				key,
				config.Instance.GetString(config.CommandImage(key)),
				config.Instance.GetString(config.CommandTag(key)),
				config.Instance.GetString(config.CommandParameter(key)),
				config.Instance.GetString(config.CommandCommand(key)),
				config.Instance.GetString(config.CommandPath(key)),
			})
		}
	}

	return data
}
