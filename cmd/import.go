package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	importCmd = &cobra.Command{
		Use:   "i",
		Short: "import tools",
		Long:  "import tools",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("import")
		},
	}
)

func init() {
	rootCmd.AddCommand(importCmd)
}
