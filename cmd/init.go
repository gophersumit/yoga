package cmd

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/gophersumit/yoga/executor"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create a new Yoga project",
	Long: `creates a new Yoga project with Angular as frontend
and Go as backend`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		appName := args[0]
		executor.Validate(appName)
		executor.InitServer(appName)
		executor.InitClient(appName)
		fmt.Println(emoji.CloudWithLightning.String() + " your project is ready!")
		fmt.Println("cd to " + appName + " and start hacking!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
