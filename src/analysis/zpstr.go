package analysis

import (
	"io"
	"os"
	"bufio"
	"fmt"
)

func Substring(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl-1+start
	}
	end = start+length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func Read(path string) string {

	f , err := os.Open(getFilePath(path))
	if err != nil {
		fmt.Printf("%v\n", err)
		//os.Exit(1)
		//panic(err)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	content := ""
	for {
		line , err := br.ReadString(byte('\n'))
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		content+=line
	}
	return content
}

func getFilePath(fileName string) string {
	basePath, err := os.Getwd()
	if err != nil {
		return fileName
	}
	filePath := basePath + "/" + fileName

	fmt.Println(filePath)
	return filePath
}

func binSearch(items []rune, item rune) bool {
	var low, mid int = 0, 0
	hight := len(items) - 1
	for low <= hight {
		mid = (low+hight)/2
		if items[mid] == item {
			return true
		} else if items[mid] > item {
			hight = mid-1
		} else {
			low = mid+1
		}
	}
	return false

}

func clearTagSuffix(tag []rune) []rune {
	end := len(tag) - 1
	for end >= 0 {
		if tag[end] == 10 || tag[end] == 13 || tag[end] == 0 {
			end--
		}else {
			break
		}
	}
	if end >= 0 {
		return tag[0:end+1]
	}else {
		return tag[0:0]
	}
}

func TrimSymbol(val []rune) []rune {
	end := len(val) - 1
	for end >= 0 {
		if binSearch(spaceSymbol, val[end]) {
			end--
		}else {
			break
		}
	}
	start := 0
	for start < end {
		if binSearch(spaceSymbol, val[start]) {
			start++
		}else {
			break
		}
	}

	if end >= 0 {
		return val[start:end+1]
	}else {
		return val[0:0]
	}
}

func buildNewTag(tag []rune) []rune {
	newTag := make([]rune, len(tag)*2)
	for i, j := 0, 0; i < len(tag); i, j = i+1, j+1 {
		if tag[i] != 58 {
			newTag[j] = 32
			j++
		}
		newTag[j] = tag[i]
	}
	return clearTagSuffix(newTag)
}
//12290:46,  .


func formatContent(content []rune) []rune {
	newContent := make([]rune, len(content))
	j := 0
	for i := 0; i < len(content); i++ {
		if item, ok := replaceMap[content[i]]; ok {
			newContent[j] = item
		}else {
			newContent[j] = content[i]
		}
		if newContent[j] == 32 && j > 0 {
			if newContent[j-1] == 32 {
				j--
			}
		}
		if newContent[j] == 10 && j > 0 {
			if newContent[j-1] == 10 {
				j--
			}
		}
		j++
	}
	return newContent
}


func readLine(text []rune , act func(int, []rune)) {
	position := 0
	row := 0
	for position < len(text) {
		start := position
		for position < len(text) && text[position] != 10 {
			position++
		}
		act(row, text[start:position])
		position++
		row++
	}
}

func RuneIndex(s []rune, sub []rune) (int, int) {
	start, i, j := 0, 0, 0

	//position, start := 0, 0
	for i < len(s) {
		for j < len(sub) && s[i] != sub[j] {
			i++
			j = 0
			start = i
		}
		if (i-start) == len(sub) {
			break
		}
		i++
		j++
	}
	fmt.Println("start,position", "=", start, ",", i)
	return start, i
}


func ContainForRune(items []rune, item []rune) bool {
	//	fmt.Println("item=",item)
	//	fmt.Println("items=",items)
	start := item[0]
	contain := false
	for i := 0; i < len(items); i++ {
		if items[i] == start {
			if len(items)-i < len(item) {
				return false
			}
			contain = true
			for j := 0; j < len(item); j++ {
				if item[j] != items[j+i] {
					contain = false
					break
				}
			}
		}
		if contain {
			break
		}
	}
	//fmt.Println("contain=",contain)
	return contain
}

func ClearHtmlTag(content []rune) []rune {
	text := make([]rune, 0)

	stack := make([]rune, 0)

	//fmt.Println("init stack=",len(stack))

	for _, i := range (content) {

		switch {
		case i == 60:
			stack = append(stack, i)

		case i == 62:
			//fmt.Println("62 then=",len(stack))
			stack = append(stack, i)
			isBr:=false
			//fmt.Println(stack,",",string(stack))
			for _,tag:=range(brTags){
				//fmt.Println(tag,",",string(tag))
				if ContainForRune(stack,tag){

					isBr=true
					break
				}
			}
			stack = append(stack[:0],stack[len(stack):]...)

			//fmt.Println("62 then=",len(stack))
			if isBr{
				text = append(text, 10)
			}else{
				text = append(text, 32)
			}
		default:
			//fmt.Println(i,",",len(stack))
			if len(stack) > 0 {
				stack = append(stack, i)

			}else {
				text = append(text, i)
			}
		}

	}
	if len(stack) > 0 {
		text = append(text, stack...)
	}

	return text
}
