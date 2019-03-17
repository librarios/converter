package main

import (
	"github.com/librarios/librc/app"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func scan(c *cli.Context) error {
	directory := c.Args().First()
	if !c.Args().Present() {
		log.Println("scan directory is not set.")
		cli.ShowCommandHelpAndExit(c, "scan", 1)
	}

	librcApp := app.NewApp()
	librcApp.Init()
	librcApp.Scan(directory)
	return nil
}

func upload(c *cli.Context) error {
	filePath := c.Args().First()
	if !c.Args().Present() {
		log.Println("upload filepath is not set")
		cli.ShowCommandHelpAndExit(c, "upload", 1)
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
			Name:    "scan",
			Aliases: []string{"s"},
			Usage:   "scan `directory`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "set output `filename`",
				},
			},
			Action: scan,
		},
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
