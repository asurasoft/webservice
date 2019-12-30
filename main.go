package main

import (
	"fmt"
	"os"
	"webserice/cmd/http"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "testCobra",
	Short: "A test demo",
	Long:  `Demo is a test appcation for print things`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("no name")
			return
		}
		if args[0] == "http" {
			http.Run()
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {
	Execute()
}
