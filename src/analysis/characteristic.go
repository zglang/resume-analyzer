package analysis

var brTags [][]rune

func initBrTags(){
	var htmlTags []string = []string{"br/>","br />","br>","p>","div>","tr>","h1>","h2>","h3>"}
	for _,tag:= range(htmlTags){
		rt:=[]rune(tag)
		brTags=append(brTags,rt)
	}
}
