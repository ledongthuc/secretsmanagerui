package main

import (
	"github.com/ledongthuc/secretsmanagerui/views"
)

func main() {
	app := new(views.App)
	err := app.Init()
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}
