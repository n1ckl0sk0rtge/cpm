package cmd

import (
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func list(_ *cobra.Command, args []string) {

	containers := viper.Get(config.Container).(map[string]interface{})

	var data [][]string

	for key, _ := range containers {
		data = append(data, []string{
			key,
			viper.Get(config.ContainerImage(key)).(string),
			viper.Get(config.ContainerParameter(key)).(string),
			viper.Get(config.ContainerCommand(key)).(string),
			viper.Get(config.ContainerPath(key)).(string),
		})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NAME", "IMAGE", "PARAMETER", "COMMAND", "PATH"})
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetColumnSeparator("")
	table.SetHeaderLine(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.Render()

}
