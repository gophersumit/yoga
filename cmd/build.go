package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/enescakir/emoji"
	"github.com/gophersumit/yoga/executor"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command for yoga
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Yoga project",
	Long: `This will build client Angular Application
as well Go application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(emoji.FlexedBiceps.String() + " building application ")
		checkIfProjectExists()
	},
}

func checkIfProjectExists() {
	_, err := os.Open("yoga.json")
	if err != nil {
		log.Fatalln(emoji.PersonFacepalming.String() + " Can only be run in existing yoga project")

	}
	buildProject()
}

func buildProject() {
	ngBuild := executor.Executor{
		CommandText:      "ng build --outputPath=static",
		CommandDirectory: "client/angular",
		BindOutput:       false,
	}
	err := ngBuild.Execute()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(emoji.NerdFace + " successfully build Angular code")
	goBuild := executor.Executor{
		CommandText: "go build ./server/cmd/web",
	}

	err = goBuild.Execute()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(emoji.NerdFace + " successfully build Go code")
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
