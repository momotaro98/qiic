package main

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"

	"github.com/urfave/cli"
)

const Version = "1.2.0"

func main() {
	app := cli.NewApp()
	app.Name = "qiic"
	app.Usage = "Get Qiita Stocked Articles and Access the Web Page"
	app.UsageText = "qiic command [arguments...]"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:    "access",
			Aliases: []string{"a", "open"},
			Usage:   "access to the Access Number Article Page",
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
			Usage:   "list local saved articles",
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
			Name:    "stock",
			Aliases: []string{"s"},
			Usage:   "update the stocked articles and Access Number",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "username, user, u",
					Usage:  "qiita username",
					EnvVar: "QIITA_USERNAME",
					Value:  "",
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

				if username == "" {
					msg := "username is required"
					fmt.Println(msg)
					return fmt.Errorf(msg)
				}

				// Fetch from API Server
				req := &UserStockRequest{
					UserName: username,
					GetRequest: GetRequest{
						Page: page,
					},
				}
				articles, _, err := GetArticles(context.Background(), req)
				if err != nil {
					fmt.Println(err)
					return err
				}

				return SaveAndRender(articles)
			},
		},
		{
			Name:    "rank",
			Aliases: []string{"r"},
			Usage:   "LGTM ranking",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "username, user, u",
					Usage:  "qiita username",
					EnvVar: "QIITA_USERNAME",
					Value:  "",
				},
				cli.StringFlag{
					Name:   "token, t",
					Usage:  "qiita api token",
					EnvVar: "QIITA_TOKEN",
					Value:  "",
				},
				cli.IntFlag{
					Name:  "page, p",
					Value: 1,
					Usage: "page number",
				},
			},
			Action: func(ctx *cli.Context) error {
				username := ctx.String("username")
				token := ctx.String("token")
				page := ctx.Int("page")

				var (
					c   = context.Background()
					req ArticlesGetRequester
				)

				if token != "" {
					c = SetToken(c, token)
					req = &ReqGetAuthenticatedUserItems{
						GetRequest{
							Page: 1,
						},
					}
				} else if username != "" {
					c = SetUserName(context.Background(), username)
					req = &ReqGetUserItems{
						GetRequest{
							Page: 1,
						},
					}
				} else {
					msg := "either token or username is required"
					fmt.Println(msg)
					return fmt.Errorf(msg)
				}

				// Fetch from API Server
				articles, err := CollectUserItems(c, req)
				if err != nil {
					fmt.Println(err)
					return err
				}

				sort.SliceStable(articles, func(i, j int) bool {
					return articles[i].LikesCount > articles[j].LikesCount
				})

				// Paging
				const ArtPerPage = 15
				start := ArtPerPage * (page - 1)
				end := start + ArtPerPage
				num := len(articles)
				if num <= start {
					articles = []*Article{}
				} else if num < end {
					articles = articles[start:num]
				} else {
					articles = articles[start:end]
				}

				return SaveAndRender(articles)
			},
		},
	}

	app.Run(os.Args)
}

func SaveAndRender(articles []*Article) error {
	var wg sync.WaitGroup

	// Save file
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := Save(articles); err != nil {
			fmt.Println("saving a file for cache failed")
		}
	}()

	// Display
	wg.Add(1)
	go func() {
		defer wg.Done()
		Render(articles)
	}()

	wg.Wait()
	return nil
}
