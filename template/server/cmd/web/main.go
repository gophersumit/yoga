package main

import "log"

func main() {
	app := NewYoga(true)
	err := app.Serve()
	log.Println("Error", err)
}
