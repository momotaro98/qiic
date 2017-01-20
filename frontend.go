package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/moznion/go-unicode-east-asian-width"
)

const (
	id_name            = "ID"
	max_id_half_len    = 10
	title_name         = "TITLE"
	max_title_half_len = 51
	tag_name           = "TAG"
	max_tag_half_len   = 21
	stock_name         = "STOCK"
	max_stock_half_len = 5
)

type Column interface {
	GetMaxHalfCharLen() int
	MakeFilledInStr(split_char string) string
	MakeCenterAlignedStr() string
	GenerateTurnedLines(art Article) []string
}

// IDColumn
type IDColumn struct {
	Name           string
	MaxHalfCharLen int
}

func (c IDColumn) GetMaxHalfCharLen() int {
	return c.MaxHalfCharLen
}

func (c IDColumn) MakeFilledInStr(split_char string) string {
	return strings.Repeat(split_char, c.MaxHalfCharLen)
}

func (c IDColumn) MakeCenterAlignedStr() string {
	return CenterAligned(c.Name, c.MaxHalfCharLen)
}

func (c IDColumn) GenerateTurnedLines(art Article) []string {
	i_str := strconv.FormatInt(art.ID, 10)
	rest_len := c.MaxHalfCharLen - len(i_str)
	stock_line := i_str + strings.Repeat(" ", rest_len)
	return []string{stock_line}
}

// TitleColumn
type TitleColumn struct {
	Name           string
	MaxHalfCharLen int
}

func (c TitleColumn) GetMaxHalfCharLen() int {
	return c.MaxHalfCharLen
}

func (c TitleColumn) MakeFilledInStr(split_char string) string {
	return strings.Repeat(split_char, c.MaxHalfCharLen)
}

func (c TitleColumn) MakeCenterAlignedStr() string {
	return CenterAligned(c.Name, c.MaxHalfCharLen)
}

func (c TitleColumn) GenerateTurnedLines(art Article) []string {
	return MakeTurnedLines(art.Title, c.MaxHalfCharLen)
}

// TagColumn
type TagColumn struct {
	Name           string
	MaxHalfCharLen int
}

func (c TagColumn) GetMaxHalfCharLen() int {
	return c.MaxHalfCharLen
}

func (c TagColumn) MakeFilledInStr(split_char string) string {
	return strings.Repeat(split_char, c.MaxHalfCharLen)
}

func (c TagColumn) MakeCenterAlignedStr() string {
	return CenterAligned(c.Name, c.MaxHalfCharLen)
}

func (c TagColumn) GenerateTurnedLines(art Article) []string {
	var tags_name_list []string
	for _, tag := range art.Tags {
		tags_name_list = append(tags_name_list, tag.Name)
	}
	tags_name := strings.Join(tags_name_list, ",")
	return MakeTurnedLines(tags_name, c.MaxHalfCharLen)
}

// StockColumn
type StockColumn struct {
	Name           string
	MaxHalfCharLen int
}

func (c StockColumn) GetMaxHalfCharLen() int {
	return c.MaxHalfCharLen
}

func (c StockColumn) MakeFilledInStr(split_char string) string {
	return strings.Repeat(split_char, c.MaxHalfCharLen)
}

func (c StockColumn) MakeCenterAlignedStr() string {
	return CenterAligned(c.Name, c.MaxHalfCharLen)
}

func (c StockColumn) GenerateTurnedLines(art Article) []string {
	i_str := strconv.Itoa(art.StockCount)
	rest_len := c.MaxHalfCharLen - len(i_str)
	stock_line := strings.Repeat(" ", rest_len) + i_str
	return []string{stock_line}
}

// Table
type Table struct {
	Columns []Column
}

func (t Table) GenerateTopLine() string {
	l_list := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		l_list[i] = c.MakeFilledInStr("─")
	}
	return "┌" + strings.Join(l_list, "┬") + "┐"
}

func (t Table) GenerateColumnnameLine() string {
	l_list := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		l_list[i] = c.MakeCenterAlignedStr()
	}
	return "|" + strings.Join(l_list, "|") + "|"
}

func (t Table) GenerateSperateLine() string {
	l_list := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		l_list[i] = c.MakeFilledInStr("─")
	}
	return "|" + strings.Join(l_list, "|") + "|"
}

// util func
func CenterAligned(str string, max int) (ret string) {
	rest_num := max - len(str)
	ret = strings.Repeat(" ", rest_num/2) + str + strings.Repeat(" ", rest_num-rest_num/2)
	return
}

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

// util func
func FindStringsMaxLen(lines_list [][]string) int {
	max_len := 1
	for _, lines := range lines_list {
		if len(lines) > max_len {
			max_len = len(lines)
		}
	}
	return max_len
}

// semi? util func
func MakeFullLines(lines []string, max_hchar int, max_lines int) []string {
	if diff_lines_len := max_lines - len(lines); diff_lines_len > 0 {
		for i := 0; i < diff_lines_len; i++ {
			lines = append(lines, strings.Repeat(" ", max_hchar))
		}
	}
	return lines
}

// func for Rendering
func GenerateArticleLines(art Article, table Table) []string {
	columns_line_list := make([][]string, len(table.Columns))
	for i, c := range table.Columns {
		columns_line_list[i] = c.GenerateTurnedLines(art)
	}

	// find max len
	maxInnerLinesLen := FindStringsMaxLen(columns_line_list)

	// make full lines
	for i, c := range table.Columns {
		columns_line_list[i] = MakeFullLines(columns_line_list[i],
			c.GetMaxHalfCharLen(), maxInnerLinesLen)
	}

	// join the columns factor to line
	artInnerLineList := make([][]string, maxInnerLinesLen)
	for _, cll := range columns_line_list {
		for j, cl := range cll {
			artInnerLineList[j] = append(artInnerLineList[j], cl)
		}
	}
	// make ret lines
	art_lines := make([]string, maxInnerLinesLen)
	for i, ails := range artInnerLineList {
		art_lines[i] = "|" + strings.Join(ails, "|") + "|"
	}
	return art_lines
}

// interface func of frontend
func Render(arts []Article) {
	// ### Change Point Start ###
	// Initialize Columns
	idColumn := IDColumn{Name: id_name, MaxHalfCharLen: max_id_half_len}
	titleColumn := TitleColumn{Name: title_name, MaxHalfCharLen: max_title_half_len}
	tagColumn := TagColumn{Name: tag_name, MaxHalfCharLen: max_tag_half_len}
	stockColumn := StockColumn{Name: stock_name, MaxHalfCharLen: max_stock_half_len}

	// make column list
	columns := []Column{idColumn, titleColumn, tagColumn, stockColumn}
	// ### Change Point End ###

	// Register Columns to Table
	table := Table{Columns: columns}

	// Print Top Line (Firt Line)
	fmt.Println(table.GenerateTopLine())
	// Print Column Name Line (Second Line)
	fmt.Println(table.GenerateColumnnameLine())

	// Print Articles
	for _, art := range arts {
		// Print Sperate Line
		fmt.Println(table.GenerateSperateLine())

		// Print Article Lines
		art_lines := GenerateArticleLines(art, table)
		for _, article_line := range art_lines {
			fmt.Println(article_line)
		}
	}

	// Print Foote
	fmt.Println(table.GenerateSperateLine())
}
