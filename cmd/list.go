package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/command"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Args:  cobra.ExactArgs(0),
	Short: "List all existing commands",
	Long: `
List all existing commands created by cpm.
`,
	Run: list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(_ *cobra.Command, _ []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NAME", "IMAGE", "TAG", "PARAMETER", "COMMAND"})
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

	commandConfFiles, err := command.List()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	var data [][]string

	for _, fileName := range commandConfFiles {

		maybeValues := command.ReadConfig(fileName, config.GetConfigProperties())
		if maybeValues == nil {
			err = fmt.Errorf("could not read/found command command config")
			fmt.Println(err)
			return nil
		}
		values := *maybeValues

		data = append(data, []string{
			values[command.Name],
			values[command.Image],
			values[command.Tag],
			values[command.Parameter],
			values[command.Commands],
		})
	}

	return data
}
