package main

import (
	"log"

	"github.com/samitghimire/botapi/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
