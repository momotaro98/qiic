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
		username := ctx.String("username")
		fmt.Println("username: ", username)
		/*
			//TEST
			if tag := ctx.String("tag"); tag == "go" {
				fmt.Println("golang pages")
			} else if tag == "python" {
				fmt.Println("python pages")
			}

			us := NewUserStockAPI(username)
			articles := us.Fetch()
			fmt.Println("articles: ", articles)
		*/
		as := NewSampleArticles()
		for _, a := range as {
			// fmt.Println(a.ID, a.User, a.Title, a.Tags)
			fmt.Println(a.ID, a.Tweet)
		}

		/*
			TODO: make runnable structs
				[X]: backend struct (has full api url, Fetch)
				[X]: Fetch Articles
				[]: Render with frontend struct
				[]: Filter Articles with tag, title etc
		*/

		/*
			// generate backend struct including full api url
			us := NewUserStockAPI(username)  // us will implement
			// fetch the url's api data and
			// store the data as original struct type(articles, tags, users...etc)
			arts := us.Fetch()

			// filter the articles according to specified tag and title name
			// arts, err := arts.Filter(tag, title)  // TODO: the arguments should be capsuled
			//if err != nil {
			//	log.Fatalf("Error!")
			//}

			// select frontend type
			fe, ok := getFrontEnd(ctx)
			if !ok {
				log.Fatalf("Not OK")
			}
			// render according to the articles data and frontend type
			fe.Render(arts)
		*/
	}

	app.Run(os.Args)
}
