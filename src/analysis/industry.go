package analysis

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

var IndustryMap map[rune][][]rune

func initIndustry() {
	IndustryMap = make(map[rune][][]rune)
	fmt.Println("行业词库")
	f , err := os.Open(getFilePath("conf/industry.txt"))
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

			putStringToMap(line, IndustryMap)

		}
	}

	fmt.Println(IndustryMap)
}
