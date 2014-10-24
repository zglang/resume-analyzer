package controller

import (
	"fmt"
	"net/http"
	"strings"
	"analysis"
	"encoding/json"
	"html/template"
)

type ResultBody struct {
	Message string
}

func ReloadController(w http.ResponseWriter, r *http.Request) {
	initForm(r)
	analysis.InitConf()
	var text ResultBody = ResultBody{"重新加载成功"}
	result, err := json.Marshal(text)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	fmt.Fprintf(w, string(result))
}

func AnalysisController(w http.ResponseWriter, r *http.Request) {
	initForm(r)
	content := r.Form["content"][0]
	resume := analysis.Analysis(content)
	result, err := json.Marshal(resume)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	w.Header().Add("Content-Type", "text/json; charset=utf-8")
	fmt.Fprintf(w, string(result))
}

func ReadResumeController(w http.ResponseWriter, r *http.Request) {
	initForm(r)

//	defer func() {
//		fmt.Println("defer func()")
//		if err:=recover();err!=nil{
//			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
//		}
//	}()

	cvfile := r.Form["cvfile"][0]
	content := analysis.Read(cvfile + ".txt")
	if r.Method == "GET" {
		fmt.Fprintf(w, content)
	}else if r.Method == "POST" {
		t, _ := template.ParseFiles("gtpl/submit.gtpl")
		t.Execute(w, content)
	}
}

func SubmitController(w http.ResponseWriter, r *http.Request) {
	initForm(r)
	content := analysis.Read("zhilian.txt")
	t, _ := template.ParseFiles("gtpl/submit.gtpl")
	t.Execute(w, content)
}

func initForm(r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}

