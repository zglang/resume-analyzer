package analysis

import (
	"regexp"
)

var unNameTags []rune = []rune{9, 10, 13, 32, 124}

func matchName(position int, text []rune, cv *Resume) int {
	for _,t := range cv.Items{
		if t.Name == TagList[0] && t.Name != "" {
			return position
		}
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
			//fmt.Println("姓名:", name)
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

var dateTags []rune = []rune{32, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 24180, 26085, 26376}

func matchDate(position int, text []rune, cv *Resume) int {
	for _,t := range cv.Items{
		if t.Name == TagList[2] && t.Name != "" {
			return position
		}
	}
	for position < len(text) {
		start := position

		i := 0
		for binSearch(dateTags, text[position]) {
			if text[position] == 32 && i == 0 {
				break
			}
			position++
			i++
		}
		if i > 0 {
			if position-start > 5 {
				cv.Items = append(cv.Items, CVItem{TagList[2], string(text[start:position])})
				break
			}
		}
		position++
	}
	return position
}

var excludeNames [][]rune = [][]rune{{20010, 20154}, {31616, 21382}, {25307, 32856}, {27714, 32844}}

