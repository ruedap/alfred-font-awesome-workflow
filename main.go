package main

import (
	"os"

	"github.com/codegangsta/cli"
)

const version string = "0.1.0"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "faw"
	app.Version = version
	app.Usage = "Font Awesome Workflow for Alfred"
	app.Author = "ruedap"
	app.Email = "ruedap@ruedap.com"
	app.Commands = commands
	return app
}
