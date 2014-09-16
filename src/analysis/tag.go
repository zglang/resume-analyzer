package analysis

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

var TagMap TMap
var GroupMap TMap
var TagList []string
var spaceSymbol []rune = []rune{0, 9, 10, 13, 32, 40, 58, 124}

type TMap map[rune][]TMapItem

type TMapItem struct {
	Body     []rune
	TagIndex int
	Type     int
}

type Tag struct {
	id      int
	name    string
	content string
	tagType int
	start   int
}

func (tmap TMap) putTag(index int, text string) string {
	tags := strings.Split(text, ",")
	tagType, _ := strconv.Atoi(tags[0])
	for i := 0; i < len(tags); i++ {
		if i == 0 {
			continue
		}
		runes := []rune(tags[i])
		tmpTag := clearTagSuffix(runes[1:])
		end := len(tmpTag) - 1
		addNewItem := true

		if tmpTag[end] == 124 {
			tmpTag = tmpTag[:end]
			addNewItem = false
			fmt.Println("含有|", tmpTag)
		}
		//
		//		}


		subItem := TMapItem{tmpTag, index, tagType}
		var additionItem TMapItem
		if addNewItem {
			additionItem = TMapItem{buildNewTag(subItem.Body), index, tagType}
		}else {
			additionItem = subItem
		}

		if _, ok := tmap[runes[0]]; ok {
			item := sort(tmap[runes[0]], subItem)
			item2 := sort(tmap[runes[0]], additionItem)

			//tmap[runes[0]] = append(tmap[runes[0]], item)

			tmap[runes[0]] = append(tmap[runes[0]], item2)
			tmap[runes[0]] = append(tmap[runes[0]], item)
			//tmap[runes[0]]= append(tmap[runes[0]],TMapItem{runes[1:],index})
		}else {
			items := make([]TMapItem, 40)
			items = []TMapItem{additionItem, subItem}
			tmap[runes[0]] = items
		}
	}
	return tags[1]
}

func (tmap TMap) putGroup(text string) {
	tags := strings.Split(text, ",")
	tagIndex, _ := strconv.Atoi(tags[0])
	for i := 0; i < len(tags); i++ {
		if i == 0 {
			continue
		}
		runes := []rune(tags[i])
		subItem := TMapItem{clearTagSuffix(runes[1:]), tagIndex, 3}
		additionItem := TMapItem{buildNewTag(subItem.Body), tagIndex, 3}
		if _, ok := tmap[runes[0]]; ok {
			item := sort(tmap[runes[0]], subItem)
			item2 := sort(tmap[runes[0]], additionItem)
			tmap[runes[0]] = append(tmap[runes[0]], item)
			tmap[runes[0]] = append(tmap[runes[0]], item2)
			//tmap[runes[0]]= append(tmap[runes[0]],TMapItem{runes[1:],index})
		}else {
			items := make([]TMapItem, 30)
			items = []TMapItem{additionItem, subItem}
			tmap[runes[0]] = items
		}
	}

}

func sort(items []TMapItem, item TMapItem) TMapItem {
	for i := 0; i < len(items); i++ {
		if len(item.Body) > len(items[i].Body) {
			var tmp TMapItem
			for j := i; j < len(items); j++ {
				tmp = items[j]
				items[j] = item
				item = tmp
			}
		}
	}
	return item
}

func initMap() {
	fmt.Println("初始化标签")
	f , err := os.Open(getFilePath("conf/tag.txt"))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	var index int
	for {
		line , err := br.ReadString(byte('\n'))
		if err == io.EOF {
			fmt.Println(err)
			break
		}else {
			tagName := TagMap.putTag(index, line)
			TagList[index] = tagName
		}
		index++
	}

	fmt.Println("TagMap=", TagMap)
}

func initGroup() {
	fmt.Println("分组")
	f , err := os.Open(getFilePath("conf/group.txt"))
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		line , err := br.ReadString(byte('\n'))
		if err == io.EOF {
			fmt.Println(err)
			break
		}else {
			GroupMap.putGroup(line)
		}
	}
}

