package main

import (
	"fmt"
	"github.com/librarios/librc/app"
	"github.com/urfave/cli"
	"log"
	"os"
	"sort"
	"time"
)

func printError(msg string) {
	fmt.Printf("ERROR: %s\n\n", msg)
}

func scan(c *cli.Context) error {
	directory := c.Args().First()
	if !c.Args().Present() {
		printError("scan directory is not set.")
		cli.ShowCommandHelpAndExit(c, "scan", 1)
	}

	librcApp := app.NewApp()
	librcApp.Init()
	librcApp.Scan(directory, &app.ScanOption{
		OutputFile: c.String("file"),
	})
	return nil
}

func upload(c *cli.Context) error {
	filePath := c.Args().First()
	if !c.Args().Present() {
		printError("upload filepath is not set")
		cli.ShowCommandHelpAndExit(c, "upload", 1)
	}

	librcApp := app.NewApp()
	librcApp.Init()
	_, err := librcApp.Upload(filePath, c.String("bucket"), c.String("object"))
	return err
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
					Name:   "file, f",
					Usage:  "set output `filename`",
					EnvVar: "LIBRC_SCAN_OUTPUT",
				},
			},
			Action: scan,
		},
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "upload `filename`",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "bucket, b",
					Usage:  "set destination `bucketName`",
					EnvVar: "LIBRC_UPLOAD_BUCKET",
				},
				cli.StringFlag{
					Name:   "object, o",
					Usage:  "set destination `objectName`",
					EnvVar: "LIBRC_UPLOAD_OBJECT",
				},
			},
			Action: upload,
		},
	}

	sort.Sort(cli.FlagsByName(cliApp.Flags))
	sort.Sort(cli.CommandsByName(cliApp.Commands))

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
