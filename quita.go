package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const version string = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "quita"
	app.Usage = "Display Qiita' Stock and Access the Web Page"
	app.UsageText = "quita command [arguments...]"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "username, u",
			Value:  "",
			Usage:  "qiita username",
			EnvVar: "QIITA_USERNAME",
		},
		cli.StringFlag{
			Name:  "tag, t",
			Value: "",
			Usage: "query stock with tag name",
		},
		cli.StringFlag{
			Name:  "title, ttl",
			Value: "",
			Usage: "query stock with title name",
		},
	}

	app.Action = func(ctx *cli.Context) {
		fmt.Println("quita application!!!")
		// be, err = getBackEnd()
		if tag := ctx.String("tag"); tag == "go" {
			fmt.Println("golang pages")
		} else if tag == "python" {
			fmt.Println("python pages")
		}
	}
	app.Run(os.Args)
}
