package main

import (
	// "fmt"
	"os"

	"github.com/urfave/cli"
)

/*
	TODO: make runnable structs
		Main Func
		[X]: backend struct (has full api url, Fetch)
		[X]: Fetch Articles
		[X]: Render with frontend struct
			[X]: To Have Structs
		[]: Access to the WebPage

		Sub Func
		[]: Filter Articles with tag, title etc
*/

const version string = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "quita"
	app.Usage = "Display Qiita' Stock and Access the Web Page"
	app.UsageText = "quita command [arguments...]"
	app.Version = version
	app.Commands = []cli.Command{
		/*
			{
				Name:    "access",
				Aliases: []string{"a"},
				Usage:   "access to the Access Number WebPage",
				Action: func(c *cli.Context) error {
					fmt.Println("added task: ", c.Args().First())
					return nil
				},
			},
		*/
		{
			Name:    "list",
			Aliases: []string{"l", "ls"},
			Usage:   "list the local saved articles",
			Action: func(c *cli.Context) error {
				articles, err := Load()
				if err != nil {
					return err
				}
				Render(articles)
				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update the stocked articles and Access Number",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "username, user",
					Value:  "",
					Usage:  "qiita username",
					EnvVar: "QIITA_USERNAME",
				},
			},
			Action: func(ctx *cli.Context) error {
				username := ctx.String("username")
				// Fetch from API Server
				user_stock := NewUserStockAPI(username)
				articles := user_stock.Fetch()
				// Save to Local File
				err := Save(articles)
				if err != nil {
					return err
				}
				// Display
				Render(articles)
				return nil
			},
		},
	}
	/*
		// Test Using app.Action
		app.Action = func(ctx *cli.Context) {
		}
	*/
	app.Run(os.Args)
}
