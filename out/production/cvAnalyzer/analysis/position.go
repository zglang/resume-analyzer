package analysis

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

var PositionMap map[rune][][]rune
var tmpPositionMap map[string]byte

func initPosition() {
	PositionMap = make(map[rune][][]rune)
	tmpPositionMap = make(map[string]byte)
	fmt.Println("职位词库")
	f , err := os.Open(getFilePath("conf/position.txt"))
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
			if _,ok:=tmpPositionMap[line];!ok{
				tmpPositionMap[line]=0
				putStringToMap(line, PositionMap)
			}
			//putPosition(line, PositionMap)
		}
	}
	tmpPositionMap=nil
	//fmt.Println("[PositionMap]======================",len(PositionMap))
	fmt.Println(PositionMap)
}


func putStringToMap(text string, posotions map[rune][][]rune) {
	runes := []rune(text)
	if len(runes)<2{
		return
	}
	subItem := clearTagSuffix(runes[1:])
	if _, ok := posotions[runes[0]]; ok {
		item := sortRuneSlice(posotions[runes[0]], subItem)
		posotions[runes[0]] = append(posotions[runes[0]], item)
	}else {
		var items [][]rune
		items = [][]rune{subItem}
		posotions[runes[0]] = items
	}
}

func sortRuneSlice(items [][]rune, item []rune) []rune {
	for i := 0; i < len(items); i++ {
		if len(item) > len(items[i]) {
			var tmp []rune
			for j := i; j < len(items); j++ {
				tmp = items[j]
				items[j] = item
				item = tmp
			}
		}
	}
	return item
}
