package analysis

import (
	"fmt"
	"regexp"
)

type GroupProcessor interface {
	Analyze(tag *Tag) CVGroup
}

type WorkExperience struct {}

var GroupProcessors []GroupProcessor

var timeRegex []string = []string{`^[\d]{2,4}[-./\s]+[\d]{1,2}\s*[-到至]{1,4}\s*[\d]{2,4}[-./\s]+[\d]{1,2}`,
	`^[\d]{2,4}[-./\s]+[\d]{1,2}\s*[-到至]{1,4}\s*(?:(?:至今)|(?:今))`,
	`^[\d一二三四五六七八九十]{1,3}[-./\s月]+[\d]{2,4}\s*[-到至]{1,4}\s*[\d一二三四五六七八九十]{1,3}[./\s月]+[\d]{2,4}`,
	`^[\d一二三四五六七八九十]{1,3}[-./\s月]+[\d]{2,4}\s*[-到至]{1,4}\s*(?:(?:至今)|(?:今)`}

func (g *WorkExperience) Analyze(tag *Tag) CVGroup {
	var group CVGroup
	group.Name = tag.name
	group.Value = tag.content
	section := make(map[string]string)
	var content string
	var key string

	readLine([]rune(tag.content), func(row int, line []rune) {
			tmp := string(line)
			for _, exp := range timeRegex {
				reg, err := regexp.Compile(exp)

				if err != nil {
					fmt.Println("错误:", err)
				}else {

					all := reg.FindAllString(tmp, -1)
					if len(all) > 0 {
						if key != "" {
							section[key] = content
							content = ""
						}
						key = all[0]
						break;
					}
				}
			}
			if row != 0 {
				content+="\n"
			}
			content+=tmp
		})
	section[key] = content

	fmt.Println(section)
	group.Items = make([][]CVItem, len(section))
	groupIndex:=0
	for k, v := range section {
		items := make([]CVItem, 4)
		items[0] = CVItem{"工作日期", k}
		runeBody := []rune(v)
		start, position := RuneIndex(runeBody, []rune(k))
		name, position := findCompany(start, position, runeBody)
		items[1] = CVItem{"公司名称", name}
		job := findKeyWord(start, position, runeBody, PositionMap)
		if job == "-" && start > 0 {
			job = findKeyWord(start, 0, runeBody, PositionMap)
		}
		items[2] = CVItem{"职位名称", job}
		industry := findKeyWord(start, position, runeBody, IndustryMap)
		items[3] = CVItem{"所属行业", industry}
		group.Items[groupIndex] = items
		groupIndex++
	}
	return group
}

func findKeyWord(start int, position int, content []rune, words map[rune][][]rune) (string) {
	for position < len(content) {
		prefix := content[position]
		if _, ok := words[prefix]; ok {
			//fmt.Println("prefix=",string(prefix))
			for _, pos := range words[prefix] {
				start := position
				match := true
				start++

				for _, c := range pos {
					if start >= len(content) {
						break
					}
					//fmt.Println(string(c),",",string(content[start]))
					if c != content[start] {
						match = false
						break
					}
					start++
				}
				if match {
					findName := []rune{prefix}
					findName = append(findName, pos...)
					//fmt.Println("=string(findName)==========================")
					//fmt.Println(string(findName))
					return string(findName)
				}
			}
		}
		position++
	}


	//
	//	newRow:=0
	//	start = position
	//	fmt.Println("=================")
	//	for position<len(content){
	//		if content[position]==10{
	//			newRow++
	//		}
	//
	//		if position<len(content) && binSearch(spaceSymbol,content[position]){
	//			aa:=string(content[start:position])
	//			fmt.Println("row,aa=",newRow,",",aa)
	//			start=position
	//			start++
	//		}
	//		position++
	//	}
	return "-"
}

func findCompany(start int, position int, content []rune) (string, int) {
	start = position
	for position < len(content) {
		if !binSearch(spaceSymbol,content[position]) {
			if content[position] == 10 {
				break
			}
			position++
		}else {
			if start == position {
				start++
				position++
			}else {
				break
			}
		}
	}
	return string(content[start:position]), position
}



func findPositionName(position int) {

}

func initGroupProcessor() {
	GroupProcessors[28] = &WorkExperience{}
}
