package main

import (
	"fmt"
	"strings"

	"github.com/moznion/go-unicode-east-asian-width"
)

const (
	max_title_half_len = 51
	max_tag_half_len   = 21
	max_stock_half_len = 5
)

// util func
func MakeTurnedLines(str string, max_len int) (t_lines []string) {
	var isJustMaxLenFlag bool
	var cur_half_len int
	var cur_line []rune
	for _, r := range str { // 'r' means rune
		if eastasianwidth.IsFullwidth(r) {
			cur_half_len += 2
		} else {
			cur_half_len++
		}
		if cur_half_len == max_len {
			isJustMaxLenFlag = true
		} else if cur_half_len > max_len {
			// Arrange to Full and Half
			if isJustMaxLenFlag == true {
				t_lines = append(t_lines, string(cur_line))
			} else {
				t_lines = append(t_lines, string(cur_line)+" ")
			}
			// Initalize stat variables
			if eastasianwidth.IsFullwidth(r) {
				cur_half_len = 2
			} else {
				cur_half_len = 1
			}
			isJustMaxLenFlag = false
			cur_line = nil
		}
		cur_line = append(cur_line, r)
	}
	var spaces string
	if rest_len := max_len - cur_half_len; rest_len > 0 {
		spaces = strings.Repeat(" ", rest_len)
	}
	t_lines = append(t_lines, string(cur_line)+spaces)
	return
}

// util
func MakeFullLines(lines []string, max_hchar int, max_lines int) []string {
	if diff_lines_len := max_lines - len(lines); diff_lines_len > 0 {
		for i := 0; i < diff_lines_len; i++ {
			lines = append(lines, strings.Repeat(" ", max_hchar))
		}
	}
	return lines
}

// util
func FindStringsMaxLen(lines_list [][]string) int {
	max_len := 1
	for _, lines := range lines_list {
		if len(lines) > max_len {
			max_len = len(lines)
		}
	}
	return max_len
}

func GenerateArticleLines(art Article) (art_lines []string) {
	// Title
	title_lines := MakeTurnedLines(art.Title, max_title_half_len)

	// Tags
	var tags_name_list []string
	for _, tag := range art.Tags {
		tags_name_list = append(tags_name_list, tag.Name)
	}
	tags_name := strings.Join(tags_name_list, ",")
	tags_lines := MakeTurnedLines(tags_name, max_tag_half_len)

	// Stock
	stock_line := fmt.Sprintf("%5d", art.StockCount) // fixed to 5 char in stock
	stock_lines := []string{stock_line}

	// find max len
	lines_list := [][]string{title_lines, tags_lines, stock_lines}
	max_lines_len := FindStringsMaxLen(lines_list)

	// make full lines
	title_lines = MakeFullLines(title_lines, max_title_half_len, max_lines_len)
	tags_lines = MakeFullLines(tags_lines, max_tag_half_len, max_lines_len)
	stock_lines = MakeFullLines(stock_lines, max_stock_half_len, max_lines_len)

	/*
		// DeBug
		for _, title := range title_lines {
			fmt.Println(title)
		}
		for _, tag := range tags_lines {
			fmt.Println(tag)
		}
		for _, stock := range stock_lines {
			fmt.Println(stock)
		}
	*/

	art_lines = title_lines
	return art_lines
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
	// var art_lines []string
	for _, art := range arts {
		spa_line := "├───────────────────────────────────────────────────┼─────────────────────┼─────┤"
		fmt.Println(spa_line)

		// art_lines = GenerateArticleLines(art)
		GenerateArticleLines(art)
		// fmt.Println(art_lines)
	}

	// Print Footer
	footer := "└───────────────────────────────────────────────────┴─────────────────────┴─────┘"
	fmt.Println(footer)
}
