package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dtf",
		Short: "dobby terraform tools",
		Long:  "dobby terraform tools",
	}
)

func initConfig() {

}

func MustExecute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
