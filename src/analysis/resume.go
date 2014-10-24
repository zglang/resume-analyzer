package analysis

import "fmt"

type CVItem struct {
	Name     string
	Value    string
}

type CVGroup struct {
	CVItem
	Items [][]CVItem
}

type Resume struct {
	Items  []CVItem
	Groups [] CVGroup
}

func Analysis(content string) Resume {
	fmt.Println("++++++++++[start]+++++++++++")
	codingBody := formatContent(ClearHtmlTag([]rune(content)))
	position := 0
	sexPosition:=0
	var tags []Tag

	for position < len(codingBody) {
		character := codingBody[position]
		if items, ok := TagMap[character]; ok {
			match := false
			var tag Tag

			if position, match = matchItems(position, codingBody, items, &tag); match {
				tags = append(tags, tag)
				if tag.id==1{
					sexPosition=position
					fmt.Println("tag.name=",tag.name,";sexPosition=",sexPosition)
				}
			}
		}
		position++
	}

	var cv Resume


	for _, tag := range tags {

		switch {
		case MaterialCorrections[tag.id] != nil:
			content := MaterialCorrections[tag.id].Correct(tag.content)
			cv.Items = append(cv.Items, CVItem{tag.name, content})
		case GroupProcessors[tag.id] !=nil:
			group := GroupProcessors[tag.id].Analyze(&tag)
			cv.Groups = append(cv.Groups, group)
		default:
			cv.Items = append(cv.Items, CVItem{tag.name, tag.content})
		}
	}
	if sexPosition>0{
		matchName(0, codingBody[:sexPosition], &cv)
	}else{
		matchName(0, codingBody, &cv)
	}
	matchMobile(0, codingBody, &cv)
	matchDate(0, codingBody[sexPosition:], &cv)
	matchEmail(0, codingBody, &cv)
	return cv
}

func matchItems(position int, text []rune, items []TMapItem, tag *Tag) (int, bool) {
	//match := true
	for _, item := range items {
		match := true
		curr := position
		for _, sub := range item.Body {
			curr++
			//fmt.Printf("[2] position:%d,coding:%d,curr=%d,sub=%d\n", position, text[curr],curr,sub)
			if curr<len(text) && sub != text[curr] {
				match = false
				break
			}
		}

		if match {
			isTag := verifyTag(position,curr,text,items)
			if isTag{
				var previous rune = 10
				if position>0{
					previous = text[position - 1:position][0]
				}
				behind := text[curr + 1:curr + 2][0]
				if !binSearch(spaceSymbol, previous) || !binSearch(spaceSymbol, behind) {
					return position, false
				}
			}

			if curr+1 < len(text) && text[curr+1] == 58{
				curr=curr+1
			}

			tag.id = item.TagIndex
			tag.name = TagList[item.TagIndex]
			tag.tagType = item.Type
			//fmt.Println(string([]rune{text[position]}))
			if content, start, currPosition, ok := TagFilters[item.Type].Extract(text, position, curr); ok {
				tag.content = content
				tag.start = start
				return currPosition, match
			}else {
				return position, false
			}
		}
	}
	return position, false
}

func matchItem(position int, text []rune, items []TMapItem) (int, bool) {

	for i := 0; i < len(items); i++ {
		item := items[i]
		match := true
		curr := position
		for _, sub := range item.Body {
			curr++
			if curr<len(text) && sub != text[curr] {
				match = false
				break
			}
		}
		//fmt.Println(match)
		if match {
			isTag := verifyTag(position,curr,text,items)



			if isTag {
				previous := text[position - 1:position][0]
				behind := text[curr + 1:curr + 2][0]
				end := text[curr:curr + 1][0]
				//fmt.Println(string(findTag),",match=",match)
//				fmt.Println(string(text[position:curr+1]))
//				fmt.Println(string(text[position:curr+2]))
//				fmt.Println(string(text[position-1:curr+1]))
				if binSearch(spaceSymbol, previous) && (binSearch(spaceSymbol, behind) || end == 58 ) {

					fmt.Println("标题标签:",string(text[position:curr+1]),"[",position,",",curr,"],", match)
					return curr, match
				}
				fmt.Println("错误标签:",string(text[position:curr+1]),"[",position,",",curr,"],", false)
				return position, false
			}else {
				if curr+1<len(text) {
					fmt.Println("分词标签:", string(text[position:curr+1]), "[", position, ",", curr, "],", false)
					return curr, match
				}else{
					return position, false
				}
			}



		}
	}

	return position, false
}

func verifyTag(position int,curr int,text []rune,items []TMapItem) bool{
	isTag:=true
	if position+1 > len(text) || curr+1 > len(text){
		return false
	}

	findTag := text[position+1:curr+1]
	//fmt.Println(string(findTag))
	if len(findTag)>0 && findTag[0] != 32 {
		tmpTag := make([]rune, len(findTag)*2)
		for i := 0; i < len(findTag); i++ {
			tmpTag[i*2] = 32
			tmpTag[i*2+1] = findTag[i]
		}
		//fmt.Println(string(tmpTag))
		for _, item := range items {
			isTag = false
			if len(item.Body) < len(tmpTag) {
				break
			}else if len(item.Body) != len(tmpTag) {
				continue
			}else {
				isTag = true
				//						fmt.Println(item.Body)
				//						fmt.Println(tmpTag)
				for j := 0; j < len(tmpTag); j++ {
					if tmpTag[j] != item.Body[j] {
						isTag = false
						break
					}
				}
			}
			if isTag{
				break
			}
		}
	}
	return isTag
}
