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
	Short: "A brief description of your command",
	Long:  `A longer description that spans`,
	Run:   list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(_ *cobra.Command, _ []string) {

	containers := config.Instance.Get(config.Container).(map[string]interface{})

	var data [][]string

	for key := range containers {
		data = append(data, []string{
			key,
			config.Instance.GetString(config.ContainerImage(key)),
			config.Instance.GetString(config.ContainerTag(key)),
			config.Instance.GetString(config.ContainerParameter(key)),
			config.Instance.GetString(config.ContainerCommand(key)),
			config.Instance.GetString(config.ContainerPath(key)),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NAME", "IMAGE", "TAG", "PARAMETER", "COMMAND", "PATH"})
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetColumnSeparator("")
	table.SetHeaderLine(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()

}
