package analysis

import "fmt"

type CVItem struct {
	Name	string
	Value	string
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
	codingBody := formatContent([]rune(content))
	position := 0
	var tags []Tag
	for position < len(codingBody) {
		character := codingBody[position]
		if items, ok := TagMap[character]; ok {
			match := false
			var tag Tag

			if position, match = matchItems(position, codingBody, items, &tag); match {
				tags = append(tags, tag)
			}
		}
		position++
	}

	var cv Resume
	matchName(0, codingBody[:150], &cv)
	matchDate(0, codingBody[0:200], &cv)
	for _, tag := range tags {

		switch {
		case MaterialCorrections[tag.id] != nil:
			content := MaterialCorrections[tag.id].Correct(tag.content)
			cv.Items = append(cv.Items, CVItem{tag.name, content})
		case GroupProcessors[tag.id] !=nil:
			group:=GroupProcessors[tag.id].Analyze(&tag)
			cv.Groups = append(cv.Groups, group)
		default:
			cv.Items = append(cv.Items, CVItem{tag.name, tag.content})
		}
	}
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
			if sub != text[curr] {
				match = false
				break
			}
		}
		if match {
			tag.id = item.TagIndex
			tag.name = TagList[item.TagIndex]
			tag.tagType = item.Type
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
			if sub != text[curr] {
				match = false
				break
			}
		}
		//fmt.Println(match)
		if match {
			previous := text[position - 1:position][0]
			behind := text[curr + 1:curr + 2][0]
			end := text[curr:curr + 1][0]
			if binSearch(spaceSymbol, previous) && (binSearch(spaceSymbol, behind) || (end == 58 || end == 65306)) {
				return curr, match
			}

		}
	}
	return position, false
}
