package executor

import (
	"fmt"
	"runtime"

	"github.com/enescakir/emoji"
)

// InitClient initialize client code
func InitClient(appName string) {
	fmt.Println(emoji.PersonInLotusPosition.String() + " Initializing Angular frontend for:" + appName)
	createAngularApp(appName)
}

func createAngularApp(appName string) {
	checkAngularCLI()
	createApp(appName)
}

func checkAngularCLI() {
	ngVersion := Executor{
		CommandText: "ng --version",
		BindOutput:  false,
	}
	err := ngVersion.Execute()
	if err != nil {
		fmt.Println(emoji.ConstructionWorker.String() + " Installing CLI for angular first")
		if runtime.GOOS == "windows" {
			ngInstall := Executor{
				CommandText: "npm install -g @angular/cli --ignore-scripts",
			}
			err = ngInstall.Execute()
			check(err)
		} else {
			ngInstall := Executor{
				CommandText: "sudo npm install -g @angular/cli --ignore-scripts",
			}
			err = ngInstall.Execute()
			check(err)
		}

	}
}
func createApp(appName string) {
	path := "./" + appName + "/client/angular"
	createCmd := `ng new clientApp --directory=` + path
	ngCreate := Executor{
		CommandText: createCmd,
		BindOutput:  true,
	}
	err := ngCreate.Execute()
	check(err)
}
