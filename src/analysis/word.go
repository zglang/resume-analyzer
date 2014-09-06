package analysis

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"strings"
)

func initWord() {
	fmt.Println("词库")
	f , err := os.Open(getFilePath("conf/word.txt"))
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
			TagMap.putWord(line)
		}
	}
	fmt.Println(TagMap)
}

func (tmap TMap) putWord(text string) {
	tags := strings.Split(text, ",")
	tagIndex, err1 := strconv.Atoi(tags[0])
	tagType, _ := strconv.Atoi(tags[1])
	if err1 != nil {
		return
	}
	for i := 0; i < len(tags); i++ {
		if i < 2 {
			continue
		}
		runes := []rune(tags[i])
		subItem := TMapItem{clearTagSuffix(runes[1:]), tagIndex, tagType}
		if _, ok := tmap[runes[0]]; ok {
			item := sort(tmap[runes[0]], subItem)
			tmap[runes[0]] = append(tmap[runes[0]], item)
			//tmap[runes[0]]= append(tmap[runes[0]],TMapItem{runes[1:],index})
		}else {
			items := make([]TMapItem, 40)
			items = []TMapItem{subItem}
			tmap[runes[0]] = items
		}
	}
}
