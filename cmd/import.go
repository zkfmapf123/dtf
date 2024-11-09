package cmd

import (
	"fmt"

	"github.com/dobby/tf/internal"
	"github.com/spf13/cobra"
)

var (
	importCmd = &cobra.Command{
		Use:   "i",
		Short: "import tools",
		Long:  "import tools",
		Run: func(cmd *cobra.Command, args []string) {
			output := internal.MustRun("terraform", "plan")
			findOutput := internal.FindImportReturnArrs(output)

			for _, o := range findOutput {
				fmt.Println(o)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(importCmd)
}
