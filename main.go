package main

import (
	"github.com/codegangsta/cli"
	"os"
)

const Version = "0.1.0"

func main() {
	app := cli.NewApp()
	app.Name = "faw"
	app.Version = Version
	app.Usage = "Font Awesome Workflow for Alfred"
	app.Author = "ruedap"
	app.Email = "ruedap@ruedap.com"
	app.Commands = Commands
	app.Run(os.Args)
}
