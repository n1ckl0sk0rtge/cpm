package cmd

import (
	"fmt"
	"github.com/n1ckl0sk0rtge/cpm/config"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the cpm version information",
	Run:   version,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func version(_ *cobra.Command, _ []string) {
	version := config.Version
	fmt.Println(version.Name)
	fmt.Println("Version:\t", version.Cpm)
	fmt.Println("Go Version:\t", version.Golang)
	fmt.Println("OS/Arch:\t", version.Os+"/"+version.Arch)
}
