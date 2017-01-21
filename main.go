package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

const version string = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "qiic"
	app.Usage = "Display Qiita Stocked Articles and Access the Web Page"
	app.UsageText = "qiic command [arguments...]"
	app.Version = version
	app.Commands = []cli.Command{
		{
			Name:    "access",
			Aliases: []string{"a", "open"},
			Usage:   "access to the Access Number WebPage",
			Action: func(ctx *cli.Context) error {
				if len(ctx.Args()) != 1 {
					return fmt.Errorf("Error! Need one Access Number argument\nUsage Example: qiic a 3")
				}
				arg_Anum, err := strconv.Atoi(ctx.Args().First())
				if err != nil {
					return fmt.Errorf("Error! Argument is required to be number\nUsage Example: qiic a 3")
				}
				articles, err := Load()
				if err != nil {
					return err
				}
				if !(0 < arg_Anum && arg_Anum <= len(articles)) {
					return fmt.Errorf("Error! Argument number is out of range of the articles Access Number\nUsage Example: qiic a 3")
				}
				target_article := articles[arg_Anum-1] // need decrement
				err = OpenBrowser(target_article.URL)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l", "ls"},
			Usage:   "list the local saved articles",
			Action: func(ctx *cli.Context) error {
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

	app.Run(os.Args)
}
