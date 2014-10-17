package analysis

import (
	"regexp"
	"fmt"
)

var unNameTags []rune = []rune{9, 10, 13, 32, 124}

func haveTag(tagIndex int, items []CVItem) bool {
	for _, t := range items {
		if t.Name == TagList[tagIndex] {
			if t.Value != "" {
				return true
			}
			break
		}
	}
	return false
}

func matchName(position int, text []rune, cv *Resume) int {
	if haveTag(0, cv.Items) {
		return position
	}

	newLine := true
	start := position
	for position < len(text) {
		i := 0
		for position < len(text) && binSearch(unNameTags, text[position]) {
			if text[position] == 10 {
				newLine = true
			}
			position++
		}
		start = position
		for position < len(text) && !binSearch(unNameTags, text[position]) {
			position++
			i++
		}
		if i >= 2 && i <= 4 && newLine {
			isName := true
			//fmt.Println(text[start:position])
			name := string(text[start:position])
			fmt.Println("姓名:", name)
			if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", name); m {
				findName := text[start:position]
				for _, v := range excludeNames {
					if ContainForRune(findName, v) {
						isName = false
						break
					}
				}
				if isName {
					cv.Items = append(cv.Items, CVItem{TagList[0], string(text[start:position])})
					break
				}
			}
		}
		newLine = false

	}
	return position
}

var dateTags []rune = []rune{32, 10,45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 24180, 26085, 26376}

func matchDate(position int, text []rune, cv *Resume) int {
	if haveTag(2, cv.Items) {
		return position
	}
	for position < len(text) {
		start := position
		i := 0
		dtag:=0
		for position < len(text) && binSearch(dateTags, text[position]) {
			if text[position] == 32 && i == 0 {
				break
			}
			if text[position]<48 || text[position]>57{
				dtag++
			}
			if i == 0 && position-1 >= 0 && binSearch(spaceSymbol, text[position-1]) {
				position++
				i++
			}
			if i == 0 {
				break
			}else {
				position++
				i++
			}

		}
		if i > 0 {
			if position-start > 5 && position-start < 12 && dtag>0{
				cv.Items = append(cv.Items, CVItem{TagList[2], string(text[start:position])})
				fmt.Println("Date=",string(text[start:position]))
				//break
			}
		}
		position++
	}
	return position
}



func matchMobile(position int, text []rune, cv *Resume) int {
	if haveTag(8, cv.Items) {
		return position
	}
	for position < len(text) {
		findNum := 0
		start := position
		if text[position] >= 48 && text[position] <= 57 {
			if position-1 >= 0 && binSearch(spaceSymbol, text[position-1]) {
				for position<len(text) && text[position] >= 32 && text[position] <= 57 {

					position++
					findNum++
					if findNum > 16 {
						break
					}
				}
			}
		}
		if findNum >= 7 && findNum <= 16 {

			if text[position-1] < 48 || text[position-1] > 57 {
				position--
			}
			if position+1<len(text){
				if text[position+1]<48 || text[position+1] > 57 {
					cv.Items = append(cv.Items, CVItem{TagList[8], string(text[start:position])})
				}
			}else{
				cv.Items = append(cv.Items, CVItem{TagList[8], string(text[start])})
			}

		}
		position++
	}
	return position
}


func matchEmail(position int, text []rune, cv *Resume) int {
	if haveTag(9, cv.Items) {
		return position
	}
	for position < len(text) {
		if text[position] == 64{

			start:=position-1
			end:=position+1
			for !binSearch(spaceSymbol, text[end]){
				end++
			}
			for !binSearch(spaceSymbol, text[start]){
				start--
			}

			cv.Items = append(cv.Items, CVItem{TagList[9], string(text[start+1:end])})
			return end
		}
		position++
	}
	return position
}

