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
		/*
			TODO: make runnable structs
				[X]: backend struct (has full api url, Fetch)
				[X]: Fetch Articles
				[X]: Render with frontend struct
					[X]: To Have Structs
				[]: Access to the WebPage
				[]: Filter Articles with tag, title etc
		*/

		// *** Main Start ***
		username := ctx.String("username")
		// tag := ctx.String("tag")
		// title := ctx.String("title")

		user_stock := NewUserStockAPI(username)
		articles := user_stock.Fetch()
		Render(articles)
		// *** Main End ***

		/* // TEST NewTestArticles
		as := NewTestArticles()
		for _, a := range as {
			// fmt.Println(a.ID, a.User, a.Title, a.Tags)
			fmt.Println(a.ID, a.Tweet)
		}
		Render(as)
		*/

		/* Future Code???
		// generate backend struct including full api url
		be := NewUserStockAPI(username)  // UserStockAPI implements Backend Interface
		// fetch the url's api data and
		// store the data as original struct type(articles, tags, users...etc)
		arts := be.Fetch()

		// filter the articles according to specified tag and title name
		arts, err := arts.Filter(tag, title)  // TODO: the arguments should be capsuled
		if err != nil {
			log.Fatalf("Error!")
		}

		// select frontend type
		fe, ok := NewCLIFontEnd(ctx)  // CLIFrontEnd implements FrontEnd Interface
		if !ok {
			log.Fatalf("Not OK")
		}
		render according to the articles data and frontend type
		fe.Render(arts)
		*/
	}

	app.Run(os.Args)
}
