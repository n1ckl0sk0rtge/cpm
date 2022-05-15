package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec COMMAND",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long:  `A longer description that `,

	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := fmt.Errorf("not enough arguments provided, need 1 got %d", len(args))
			fmt.Println(err)
			os.Exit(1)
		}
	},

	Run: execute,
}

func init() {
	rootCmd.AddCommand(execCmd)
}

func execute(cmd *cobra.Command, args []string) {
	commandName := args[0]

	viper.Get(config.Container + "." + commandName)

}
