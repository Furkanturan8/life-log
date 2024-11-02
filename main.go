package main

import (
	"github.com/bmdavis419/the-better-backend/pkg/app"
)

func main() {
	// setup and run app
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}
}
