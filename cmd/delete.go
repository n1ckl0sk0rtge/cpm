/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
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
	err := os.Remove(config.Instance.GetString(config.CommandPath(name)))
	if err != nil {
		fmt.Println(err)
		return
	}

	// remove entry in config file
	commands := config.Instance.Get(config.Container).(map[string]interface{})
	delete(commands, name)
	if len(commands) == 0 {
		config.Instance.Set(config.Container, nil)
	}
	_ = config.Instance.WriteConfig()
}
