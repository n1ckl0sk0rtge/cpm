package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/spf13/cobra"
	"os"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete COMMAND",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines`,
	Run:   deletion,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func deletion(_ *cobra.Command, args []string) {
	name := args[0]

	// remove alias file
	err := os.Remove(config.Instance.GetString(config.ExecPath) + name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// remove env file
	err = os.Remove(config.GetConfigProperties().Dir + name)
	if err != nil {
		fmt.Println(err)
		return
	}
}
