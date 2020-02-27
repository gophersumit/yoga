package executor

import "log"

import "os"

import "github.com/enescakir/emoji"

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Validate if folder is not already created
func Validate(appName string) {
	if _, err := os.Stat(appName); os.IsNotExist(err) {
		return
	}
	log.Fatalln(emoji.SkullAndCrossbones.String() + "  folder '" + appName + "' already exits!")
}
