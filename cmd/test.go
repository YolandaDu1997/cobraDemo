package cmd

import (
	"cobraDemo/pkg"

	"github.com/spf13/cobra"
)

var para string
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "This is the cobra sub command to describe your parameters",
	Long:  `The cobra demo needs a name and an age`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(name) == 0 {
			cmd.Help()
			return
		}

		pkg.ShowSub(para)
	},
}

func init() {
	TestCmd.Flags().StringVarP(&para, "para", "p", "descrip", "description of subcmd")
	RootCmd.AddCommand(TestCmd)
}
