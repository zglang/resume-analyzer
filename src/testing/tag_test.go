package testing

import (
	"testing"
	"analysis"
	"fmt"
)

func TestAnalyzer(t *testing.T) {
	analysis.InitConf()
	content:=analysis.Read("51.txt")

	fmt.Println("test start")
	for i := 0; i < 1000; i++ {
		analysis.Analysis(content)
		fmt.Println(i)
	}

	t.Error("ok end")
}

func BenchmarkAdd1(b *testing.B){

	b.StopTimer()
	analysis.InitConf()
	content:=analysis.Read("51.txt")
	b.StartTimer()
	fmt.Println("test start")
	for i := 0; i < b.N; i++ {
		analysis.Analysis(content)
	}
}


func TestInitConfig(t *testing.T){

	fmt.Println("test start")
	//analysis.InitConf()

	analysis.Read("51.txt")
	fmt.Println("end ")

	//t.Error("ok end")
}

