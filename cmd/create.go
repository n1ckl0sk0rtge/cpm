/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create NAME IMAGE",
	Args:  cobra.ExactArgs(2),
	Short: "A brief description of your command",
	Long: `
Create a new command.

 NAME is the name of the new command, which you will use to execute the command. 

 IMAGE is the OCI image o which the command is based on.
`,

	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			err := fmt.Errorf("not enough arguments provided, need 2 got %d", len(args))
			fmt.Println(err)
			os.Exit(1)
		}
	},

	Run: create,
}

type entity struct {
	Name      string
	Parameter string
	Image     string
	Command   string
	Runtime   string
}

var Entity entity

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	createCmd.Flags().StringVarP(&Entity.Parameter, "parameter", "p", "",
		"set parameter for running the container")

	createCmd.Flags().StringVarP(&Entity.Image, "command", "c", "",
		"set the command that should be executed in the container")

	createCmd.Flags().StringVarP(&Entity.Runtime, "runtime", "r", "",
		"provide container runtime, otherwise the value from the config will be used.")
}

func create(cmd *cobra.Command, args []string) {

}
