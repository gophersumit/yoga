package executor

import (
	"fmt"
	"github.com/enescakir/emoji"
	"os"
)

// InitServer - initialize server side code
func InitServer(appName string) {
	fmt.Println(emoji.PersonInLotusPosition.String() + " Initializing Go Backend for:" + appName)
	generate(appName)
}

func generate(projectName string) {
	for _, asset := range AssetNames() {
		assetInfo, err := AssetInfo(asset)
		check(err)
		if !assetInfo.IsDir() {
			RestoreAsset("", assetInfo.Name())
		}
	}
	os.Rename("template", projectName)
}
