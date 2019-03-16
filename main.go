package main

import (
	"github.com/librarios/librc/app"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func upload(c *cli.Context) error {
	filePath := c.Args().First()
	if !c.Args().Present() {
		return cli.NewExitError("upload filepath is not set", 1)
	}

	librcApp := app.NewApp()
	librcApp.Init()
	librcApp.Upload(filePath, "", "")
	return nil
}

func main() {
	cliApp := cli.NewApp()
	cliApp.Name = "librc"
	cliApp.Version = "0.0.1"
	cliApp.Compiled = time.Now()
	cliApp.Authors = []cli.Author{
		{Name: "Lechuck Roh"},
	}
	cliApp.Copyright = "(c) 2019 Lechuck Roh"
	cliApp.Usage = "librarios CLI"
	cliApp.Commands = []cli.Command{
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "upload `filepath`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bucket, b",
					Usage: "set destination `bucketName`",
				},
			},
			Action: upload,
		},
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
