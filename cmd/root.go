package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rummage",
	Short: "Using rummage",
	Long:  "Rummage is a tool used to navigate docker registry V2",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use -h to see available options")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
