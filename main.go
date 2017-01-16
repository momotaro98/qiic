package main

import (
	// "fmt"
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
		/*TEST
		username := ctx.String("username")
		fmt.Println("username: ", username)
		if tag := ctx.String("tag"); tag == "go" {
			fmt.Println("golang pages")
		} else if tag == "python" {
			fmt.Println("python pages")
		}
		*/

		/*
			TODO: make runnable structs
				[]: backend struct (has full api url, Fetch)
				[]: stock struct (future -> tag struct, user struct)
				[]: frontend struct (has Render)
		*/

		/*
			// generate backend struct including full api url
			be, err := genBackEnd(ctx)
			if err != nil {
				log.Fatalf("Error!")
			}
			// fetch the url's api data and
			// store the data as original struct type(articles, tags, users...etc)
			r := be.Fetch()
			// filter the articles according to specified tag and title name
			err := r.filter(tag, title)
			if err != nil {
				log.Fatalf("Error!")
			}

			// select frontend type
			fe, ok := getFrontEnd(ctx)
			if !ok {
				log.Fatalf("Not OK")
			}
			// render according to the r data and frontend type
			fe.Render(r)
		*/
	}

	app.Run(os.Args)
}
