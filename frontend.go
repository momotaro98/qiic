package main

import (
	"fmt"
	"strconv"
	"strings"

	eastasianwidth "github.com/moznion/go-unicode-east-asian-width"
)

const (
	accessNumName       = "A No"
	maxAccessNumHalfLen = 4
	titleName           = "TITLE"
	maxTitleHalfLen     = 51
	tagName             = "TAG"
	maxTagHalfLen       = 21
	likeName            = "LIKE"
	maxLikeHalfLen      = 5
)

// Column is struct of Column
type Column interface {
	GetMaxHalfCharLen() int
	MakeFilledInStr(splitChar string) string
	MakeCenterAlignedStr() string
	GenerateTurnedLines(artIndex int, art *Article) []string
}

// AccessNumColumn is struct
type AccessNumColumn struct {
	Name           string
	MaxHalfCharLen int
}

// GetMaxHalfCharLen is func.
func (c AccessNumColumn) GetMaxHalfCharLen() int {
	return c.MaxHalfCharLen
}

// MakeFilledInStr is func.
func (c AccessNumColumn) MakeFilledInStr(splitChar string) string {
	return strings.Repeat(splitChar, c.MaxHalfCharLen)
}

// MakeCenterAlignedStr is func.
func (c AccessNumColumn) MakeCenterAlignedStr() string {
	return CenterAligned(c.Name, c.MaxHalfCharLen)
}

// GenerateTurnedLines is func.
func (c AccessNumColumn) GenerateTurnedLines(artIndex int, art *Article) []string {
	iStr := strconv.Itoa(artIndex + 1) // On Display, Finally artIndex Increments
	restLen := c.MaxHalfCharLen - len(iStr)
	likeLine := iStr + strings.Repeat(" ", restLen)
	return []string{likeLine}
}

// TitleColumn is struct.
type TitleColumn struct {
	Name           string
	MaxHalfCharLen int
}

// GetMaxHalfCharLen is a func.
func (c TitleColumn) GetMaxHalfCharLen() int {
	return c.MaxHalfCharLen
}

// MakeFilledInStr is a func.
func (c TitleColumn) MakeFilledInStr(splitChar string) string {
	return strings.Repeat(splitChar, c.MaxHalfCharLen)
}

// MakeCenterAlignedStr is a func.
func (c TitleColumn) MakeCenterAlignedStr() string {
	return CenterAligned(c.Name, c.MaxHalfCharLen)
}

// GenerateTurnedLines is a func.
func (c TitleColumn) GenerateTurnedLines(artIndex int, art *Article) []string {
	return MakeTurnedLines(art.Title, c.MaxHalfCharLen)
}

// TagColumn is struct.
type TagColumn struct {
	Name           string
	MaxHalfCharLen int
}

// GetMaxHalfCharLen is a func.
func (c TagColumn) GetMaxHalfCharLen() int {
	return c.MaxHalfCharLen
}

// MakeFilledInStr is a func.
func (c TagColumn) MakeFilledInStr(splitChar string) string {
	return strings.Repeat(splitChar, c.MaxHalfCharLen)
}

// MakeCenterAlignedStr is a func.
func (c TagColumn) MakeCenterAlignedStr() string {
	return CenterAligned(c.Name, c.MaxHalfCharLen)
}

// GenerateTurnedLines in a func.
func (c TagColumn) GenerateTurnedLines(artIndex int, art *Article) []string {
	var tagsNameList []string
	for _, tag := range art.Tags {
		tagsNameList = append(tagsNameList, tag.Name)
	}
	tagsName := strings.Join(tagsNameList, ",")
	return MakeTurnedLines(tagsName, c.MaxHalfCharLen)
}

// LikesColumn is a struct.
type LikesColumn struct {
	Name           string
	MaxHalfCharLen int
}

// GetMaxHalfCharLen is func.
func (c LikesColumn) GetMaxHalfCharLen() int {
	return c.MaxHalfCharLen
}

// MakeFilledInStr is a func.
func (c LikesColumn) MakeFilledInStr(splitChar string) string {
	return strings.Repeat(splitChar, c.MaxHalfCharLen)
}

// MakeCenterAlignedStr is a func.
func (c LikesColumn) MakeCenterAlignedStr() string {
	return CenterAligned(c.Name, c.MaxHalfCharLen)
}

// GenerateTurnedLines is a func.
func (c LikesColumn) GenerateTurnedLines(artIndex int, art *Article) []string {
	iStr := strconv.Itoa(art.LikesCount)
	restLen := c.MaxHalfCharLen - len(iStr)
	likeLine := strings.Repeat(" ", restLen) + iStr
	return []string{likeLine}
}

// Table is struct of table
type Table struct {
	Columns []Column
}

// GenerateTopLine is func
func (t Table) GenerateTopLine() string {
	lList := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		lList[i] = c.MakeFilledInStr("─")
	}
	return "┌" + strings.Join(lList, "┬") + "┐"
}

// GenerateColumnnameLine is func
func (t Table) GenerateColumnnameLine() string {
	lList := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		lList[i] = c.MakeCenterAlignedStr()
	}
	return "|" + strings.Join(lList, "|") + "|"
}

// GenerateSperateLine is func
func (t Table) GenerateSperateLine() string {
	lList := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		lList[i] = c.MakeFilledInStr("─")
	}
	return "|" + strings.Join(lList, "|") + "|"
}

// CenterAligned is a util func
func CenterAligned(str string, max int) (ret string) {
	restNum := max - len(str)
	ret = strings.Repeat(" ", restNum/2) + str + strings.Repeat(" ", restNum-restNum/2)
	return
}

// MakeTurnedLines is a util func
func MakeTurnedLines(str string, maxLen int) (tLines []string) {
	var isJustMaxLenFlag bool
	var curHalfLen int
	var curLine []rune
	for _, r := range str { // 'r' means rune
		if eastasianwidth.IsFullwidth(r) {
			curHalfLen += 2
		} else {
			curHalfLen++
		}
		if curHalfLen == maxLen {
			isJustMaxLenFlag = true
		} else if curHalfLen > maxLen {
			// Arrange to Full and Half
			if isJustMaxLenFlag == true {
				tLines = append(tLines, string(curLine))
			} else {
				tLines = append(tLines, string(curLine)+" ")
			}
			// Initalize stat variables
			if eastasianwidth.IsFullwidth(r) {
				curHalfLen = 2
			} else {
				curHalfLen = 1
			}
			isJustMaxLenFlag = false
			curLine = nil
		}
		curLine = append(curLine, r)
	}
	var spaces string
	if restLen := maxLen - curHalfLen; restLen > 0 {
		spaces = strings.Repeat(" ", restLen)
	}
	tLines = append(tLines, string(curLine)+spaces)
	return
}

// FindStringsMaxLen is a util func
func FindStringsMaxLen(linesList [][]string) int {
	maxLen := 1
	for _, lines := range linesList {
		if len(lines) > maxLen {
			maxLen = len(lines)
		}
	}
	return maxLen
}

// MakeFullLines is a semi? util func
func MakeFullLines(lines []string, maxHchar int, maxLines int) []string {
	if diffLinesLen := maxLines - len(lines); diffLinesLen > 0 {
		for i := 0; i < diffLinesLen; i++ {
			lines = append(lines, strings.Repeat(" ", maxHchar))
		}
	}
	return lines
}

// GenerateArticleLines is func for Rendering
func GenerateArticleLines(artIndex int, art *Article, table Table) []string {
	columnsLineList := make([][]string, len(table.Columns))
	for i, c := range table.Columns {
		columnsLineList[i] = c.GenerateTurnedLines(artIndex, art)
	}

	// find max len
	maxInnerLinesLen := FindStringsMaxLen(columnsLineList)

	// make full lines
	for i, c := range table.Columns {
		columnsLineList[i] = MakeFullLines(columnsLineList[i],
			c.GetMaxHalfCharLen(), maxInnerLinesLen)
	}

	// join the columns factor to line
	artInnerLineList := make([][]string, maxInnerLinesLen)
	for _, cll := range columnsLineList {
		for j, cl := range cll {
			artInnerLineList[j] = append(artInnerLineList[j], cl)
		}
	}
	// make ret lines
	artLines := make([]string, maxInnerLinesLen)
	for i, ails := range artInnerLineList {
		artLines[i] = "|" + strings.Join(ails, "|") + "|"
	}
	return artLines
}

// Render is interface func of frontend
func Render(arts []*Article) {
	// ### Change Point Start ###
	// Initialize Columns
	accessNumColumn := AccessNumColumn{Name: accessNumName,
		MaxHalfCharLen: maxAccessNumHalfLen}
	titleColumn := TitleColumn{Name: titleName,
		MaxHalfCharLen: maxTitleHalfLen}
	tagColumn := TagColumn{Name: tagName,
		MaxHalfCharLen: maxTagHalfLen}
	likeColumn := LikesColumn{Name: likeName,
		MaxHalfCharLen: maxLikeHalfLen}

	// make column list
	columns := []Column{accessNumColumn, titleColumn, tagColumn, likeColumn}
	// ### Change Point End ###

	// Register Columns to Table
	table := Table{Columns: columns}

	// Print Top Line (Firt Line)
	fmt.Println(table.GenerateTopLine())
	// Print Column Name Line (Second Line)
	fmt.Println(table.GenerateColumnnameLine())

	// Print Articles
	for artIndex, art := range arts {
		// Print Sperate Line
		fmt.Println(table.GenerateSperateLine())

		// Print Article Lines
		artLines := GenerateArticleLines(artIndex, art, table)
		for _, articleLine := range artLines {
			fmt.Println(articleLine)
		}
	}

	// Print Foote
	fmt.Println(table.GenerateSperateLine())
}
