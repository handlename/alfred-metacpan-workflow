package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	wf "github.com/handlename/alfred-metacpan-workflow"
)

var logger = log.New(os.Stderr, "DEBUG: ", log.Ltime|log.Lshortfile)

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		logger.Println(v...)
	}
}

func debugf(format string, v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		logger.Printf(format, v...)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "Alfred Metacpan Workflow"
	app.Version = "0.0.4"
	app.Usage = ""
	app.Author = "handlename"
	app.Email = "nagata{at}handlena.me"
	app.Action = doMain
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "query",
			Value: "",
			Usage: "search query",
		},
	}
	app.Run(os.Args)
}

func doMain(c *cli.Context) {
	q := c.String("query")

	if len(q) < 3 {
		return
	}

	fmt.Println(wf.SearchModule(q))
}
