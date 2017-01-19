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
				t_lines = append(t_lines, string(cur_line)+"|")
			} else {
				t_lines = append(t_lines, string(cur_line)+" |")
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
	t_lines = append(t_lines, string(cur_line)+spaces+"|")
	return
}

func generateArticleLines(art Article) (art_lines []string) {
	// Title
	title_lines := MakeTurnedLines(art.Title, max_title_half_len)

	// Tags
	var tags_name_list []string
	for _, tag := range art.Tags {
		tags_name_list = append(tags_name_list, tag.Name)
	}
	tags_name := strings.Join(tags_name_list, ",")
	tags_lines := MakeTurnedLines(tags_name, max_tag_half_len)
	for _, line := range tags_lines {
		fmt.Println(line)
	}

	// Like
	// like_lines := art.StockCount

	// TODO: title_lines, tags_lines, and line_lines must be same length slice
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

		// art_lines = generateArticleLines(art)
		generateArticleLines(art)
		// fmt.Println(art_lines)
	}

	// Print Footer
	footer := "└───────────────────────────────────────────────────┴─────────────────────┴─────┘"
	fmt.Println(footer)
}
