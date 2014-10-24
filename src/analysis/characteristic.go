package analysis

var brTags [][]rune

func initBrTags() {
	var htmlTags []string = []string{"br/>", "br />", "br>", "<p", "<div", "<tr", "<td", "<h1", "<h2", "p>", "div>", "tr>", "td>", "h1>", "h2>"}
	for _, tag := range (htmlTags) {
		rt := []rune(tag)
		brTags = append(brTags, rt)
	}
}

var replaceMap map[rune]rune = map[rune]rune{
	65288:40, 65289:41,
	12288:32, 65306:58, 65372:124,
	11:10, 160:32, 8226:47, 7:10, 12:10, 13:10, 8211:45, 8212:45, 20:10, 21:10, 37:65285}

func initReplaceMap() {
	upper := "ABCDEFGHIJKLMNOPQRSTYWLXYZ"
	letter1 := []rune(upper)
	for i := 0; i < len(letter1); i++ {
		replaceMap[letter1[i]] = letter1[i]+32
	}
}

var spaceSymbol []rune = []rune{0, 9, 10, 13, 32, 40, 58, 63, 124}

var excludeNames [][]rune

// = [][]rune{{20010, 20154}, {31616, 21382}, {25307, 32856}, {27714, 32844}}

func initNotNameTag() {
	var unNameTags []string = []string{"招聘", "简历", "姓名", "已出国", "性别", "个人", "求职", "邮箱", "智联", "应聘", "窗体", "标签"}
	for _, tag := range (unNameTags) {
		un := []rune(tag)
		excludeNames = append(excludeNames, un)
	}

}
