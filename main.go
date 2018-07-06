package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

const version string = "1.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "qiic"
	app.Usage = "Get Qiita Stocked Articles and Access the Web Page"
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
				argAnum, err := strconv.Atoi(ctx.Args().First())
				if err != nil {
					return fmt.Errorf("Error! Argument is required to be number\nUsage Example: qiic a 3")
				}
				articles, err := Load()
				if err != nil {
					return err
				}
				if !(0 < argAnum && argAnum <= len(articles)) {
					return fmt.Errorf("Error! The argument number is out of range of the articles Access Number\nUsage Example: qiic a 3\nCheck Articles Number with\nqiic u\n  or\nqiic l")
				}
				targetArticle := articles[argAnum-1] // need decrement
				err = OpenBrowser(targetArticle.URL)
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
					Usage:  "qiita username",
					EnvVar: "QIITA_USERNAME",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 1,
					Usage: "page number",
				},
			},
			Action: func(ctx *cli.Context) error {
				username := ctx.String("username")
				page := ctx.Int("page")
				// Fetch from API Server
				userStock := NewUserStockAPI(username, page)
				articles, err := userStock.Fetch()
				if err != nil {
					return err
				}
				// Save to Local File
				err = Save(articles)
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
