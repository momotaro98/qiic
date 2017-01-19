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

func GenerateArticleLines(art Article) []string {
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

	// make ret lines
	art_lines := make([]string, len(title_lines))
	for i := 0; i < len(title_lines); i++ {
		tit_line := title_lines[i]
		tag_line := tags_lines[i]
		sto_line := stock_lines[i]
		a_lines_list := []string{tit_line, tag_line, sto_line}
		art_lines[i] = "|" + strings.Join(a_lines_list, "|") + "|"
	}
	return art_lines
}

func GenerateTopBar() (ret string) {
	title_line := strings.Repeat("─", max_title_half_len)
	tags_line := strings.Repeat("─", max_tag_half_len)
	stock_line := strings.Repeat("─", max_stock_half_len)
	line_list := []string{title_line, tags_line, stock_line}
	ret = "┌" + strings.Join(line_list, "┬") + "┐"
	return
}

// util
func CenterAligned(str string, max int) (ret string) {
	rest_num := max - len(str)
	ret = strings.Repeat(" ", rest_num/2) + str + strings.Repeat(" ", rest_num-rest_num/2)
	return
}

func GenerateColumnBar() (ret string) {
	title_line := CenterAligned("TITLE", max_title_half_len)
	tags_line := CenterAligned("TAG", max_tag_half_len)
	stock_line := CenterAligned("STOCK", max_stock_half_len)
	line_list := []string{title_line, tags_line, stock_line}
	ret = "|" + strings.Join(line_list, "|") + "|"
	return
}

func GenerateSeperateBar() (ret string) {
	title_line := strings.Repeat("─", max_title_half_len)
	tags_line := strings.Repeat("─", max_tag_half_len)
	stock_line := strings.Repeat("─", max_stock_half_len)
	line_list := []string{title_line, tags_line, stock_line}
	ret = "|" + strings.Join(line_list, "|") + "|"
	return
}

func Render(arts []Article) {
	// Print Top Header
	fmt.Println(GenerateTopBar())

	// Print Column Bar
	fmt.Println(GenerateColumnBar())

	// Print Articles
	for _, art := range arts {
		// Print Sperate Line
		fmt.Println(GenerateSeperateBar())

		// Print Article Lines
		art_lines := GenerateArticleLines(art)
		for _, article_line := range art_lines {
			fmt.Println(article_line)
		}
	}
	fmt.Println(GenerateSeperateBar())
	// Now No Footer
}
