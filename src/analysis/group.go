package analysis

import (
	"fmt"
	"regexp"
)

type GroupProcessor interface {
	Analyze(tag *Tag) CVGroup
}


type WorkExperience struct {}
type Purpose struct {}

var GroupProcessors []GroupProcessor

var timeRegex []string = []string{`^\s*[\d]{2,4}[-./\s年]+[\d]{1,2}月?\s*[-到至]{1,4}\s*[\d]{2,4}[-./\s年]+[\d]{1,2}月?`,
	`^\s*[\d]{2,4}[-./\s年]+[\d]{1,2}月?\s*[-到至]{1,4}\s*(?:(?:至今)|(?:今))`,
	`^\s*[\d]{2,4}[-./\s年]+[\d]{1,2}[-./\s月]+[\d]{1,2}日?\s*[-到至]{1,4}\s*[\d]{2,4}[-./\s年]+[\d]{1,2}[-./\s月]+[\d]{1,2}日?`,
	`^\s*[\d]{2,4}[-./\s年]+[\d]{1,2}[-./\s月]+[\d]{1,2}日?\s*[-到至]{1,4}\s*(?:(?:至今)|(?:今))`,
	`^\s*[\d一二三四五六七八九十]{1,3}[-./\s月]+[\d]{2,4}\s*[-到至]{1,4}\s*[\d一二三四五六七八九十]{1,3}[./\s月]+[\d]{2,4}`,
	`^\s*[\d一二三四五六七八九十]{1,3}[-./\s月]+[\d]{2,4}\s*[-到至]{1,4}\s*(?:(?:至今)|(?:今))`,
	`^\s*[\d]{4}\s*[-到至]{1,4}\s*[\d]{4}年`,
	`^\s*[\d]{4}\s*[-到至]{1,4}\s*(?:(?:至今)|(?:今))`}

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

	//fmt.Println(section)
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
	return "-"
}

var unCompanySymbol []rune = []rune{0, 9, 10, 32}

func findCompany(start int, position int, content []rune) (string, int) {
	start = position
	for position < len(content) {
		if !binSearch(unCompanySymbol,content[position]) {
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
	if position - start==1{
		return findCompany(position,position,content)
	}
	return string(content[start:position]), position
}

func (g *Purpose) Analyze(tag *Tag) CVGroup {
	var group CVGroup
	group.Name=tag.name
	group.Value = tag.content
	group.Items = make([][]CVItem, 1)
	codingBody:=[]rune(tag.content)
	position:=0

	for position < len(codingBody) {
		character := codingBody[position]
		if items, ok := TagMap[character]; ok {
			match := false
			var subTag Tag

			if position, match = matchItems(position, codingBody, items, &subTag); match {
				cvitem := CVItem{subTag.name, subTag.content}

				group.Items[0] = append(group.Items[0], cvitem)
			}
		}
		position++
	}
	//group.Items = make([][]CVItem, len(subTags))



	return group
}

func initGroupProcessor() {
	GroupProcessors[28] = &WorkExperience{}
	GroupProcessors[30] = &Purpose{}
}
