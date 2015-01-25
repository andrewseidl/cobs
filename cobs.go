package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/gophergala/cobs/backend"
	"github.com/gophergala/cobs/builder"
)

func main() {
	app := cli.NewApp()
	app.Name = "CoBS"
	app.Usage = "Container Build Service"
	app.Commands = []cli.Command{
		{
			Name: "builder",
			Action: func(c *cli.Context) {
				fmt.Println("run builder")
				builder.Run()
			},
		},
		{
			Name: "backend",
			Action: func(c *cli.Context) {
				fmt.Println("run backend")
				backend.Run()
			},
		},
	}
	app.Run(os.Args)
}
