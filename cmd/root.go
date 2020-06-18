package cmd

import (
	"cobraDemo/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var name string
var age int
var RootCmd = &cobra.Command{
	Use:   "cobraDemo",
	Short: "This is the cobra demo to display a self-intro",
	Long:  `The cobra demo needs a name and an age`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(name) == 0 {
			cmd.Help()
			return
		}

		pkg.Show(name, age)
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().StringVarP(&name, "name", "n", "<Your name>", "person's name")
	RootCmd.Flags().IntVarP(&age, "age", "a", 0, "person's age")
}
