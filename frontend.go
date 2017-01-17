package main

import (
	"fmt"
	"strconv"
)

func generateArticleLine(art Article) string {
	var title string
	var tags string
	var stock_num string

	// Title
	// if title = art.Title; title
	title = fmt.Sprintf("%-51s", art.Title)

	// Tags
	for _, tag := range art.Tags {
		tags = tags + ", " + tag.Name
	}
	tags = fmt.Sprintf("%-21s", tags[2:])

	// stock_num
	stock_num = fmt.Sprintf("%5s", strconv.Itoa(art.StockCount))

	art_line := "|" + title + "|" + tags + "|" + stock_num + "|"
	return art_line
}

func Render(arts []Article) {
	// Print Header
	// Title: 51 strings, Tag: 21 strings, stock_num: 5 strings
	header := append([]string{
		"┌───────────────────────────────────────────────────┬─────────────────────┬─────┐",
		"|                       TITLE                       |         TAG         |Stock|",
	})
	for _, val := range header {
		fmt.Println(val)
	}

	// Print Articles
	var art_line string
	for _, art := range arts {
		spa_line := "├───────────────────────────────────────────────────┼─────────────────────┼─────┤"
		fmt.Println(spa_line)
		art_line = generateArticleLine(art)
		fmt.Println(art_line)
	}

	// Print Footer
	footer := "└───────────────────────────────────────────────────┴─────────────────────┴─────┘"
	fmt.Println(footer)
}
