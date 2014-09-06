package analysis

import (
	"fmt"
)

type IFilter interface {
	Extract(text []rune, start int, curr int) (string, int, int, bool)
}

var TagFilters []IFilter

func initFilter() {
	TagFilters[0] = &Filter{}
	TagFilters[1] = &WordFilter{}
	TagFilters[2] = &ExpFilter{}
	TagFilters[3] = &GroupFilter{}

}

type Filter struct{}
type WordFilter struct {}
type ExpFilter struct {}
type GroupFilter struct {}
type SexFilter struct {}
type WorkExpFilter struct {}

func (filter *Filter) Extract(text []rune, start int, curr int) (string, int, int, bool) {

	start = curr+1
	position := start
	for position < len(text) {
		character := text[position]
		if items, ok := TagMap[character]; ok {
			//fmt.Println(string([]rune{character}))
			//这里可能匹配到其他前缀，所以需要一致匹配
			if _, yy := matchItem(position, text, items); yy {
				//				fmt.Println(string(text[start:position]))
				//				fmt.Println(string(ClearValueSuffix(text[start:position])))
				return string(TrimSymbol(text[start:position])), start, position-1, true
			}
		}
		position++
	}
	fmt.Println("没有匹配到标签，就将最后内容全部作为标签内容")
	return string(text[start:]), start, len(text)-1, true
}

func (filter *WordFilter) Extract(text []rune, start int, curr int) (string, int, int, bool) {

	previous := text[start - 1:start][0]
	behind := text[curr + 1:curr + 2][0]
	if binSearch(spaceSymbol, previous) && binSearch(spaceSymbol, behind) {
		return string(TrimSymbol(text[start:curr + 1])), start, curr+1, true
	}else {
		return "", start, curr, false
	}

	return "", start, curr, false
}

func (filter *ExpFilter) Extract(text []rune, start int, curr int) (string, int, int, bool) {
	position := start - 1
	for !binSearch(spaceSymbol, text[position]) && position > 0 {
		position--
	}
	exp := string(TrimSymbol(text[position + 1:curr+1]))

	return exp, start, curr+1, true
}

func (filter *GroupFilter) Extract(text []rune, start int, curr int) (string, int, int, bool) {

	if len(text)<curr+1|| len(text)<curr +2{
		return "", start, curr, false
	}
	previous := text[start - 1:start][0]
	behind := text[curr + 1:curr + 2][0]
	if binSearch(spaceSymbol, previous) && binSearch(spaceSymbol, behind) {
		position := curr + 1
		start = position
		for position < len(text) {
			character := text[position]
			if items, ok := GroupMap[character]; ok {
				if _, ok := matchItem(position, text, items); ok {
					return string(TrimSymbol(text[start:position])), start, position-1, true
				}
			}
			position++
		}
		return string(TrimSymbol(text[start:])), start, len(text)-1, true
	}else {
		return "", start, curr, false
	}
}







