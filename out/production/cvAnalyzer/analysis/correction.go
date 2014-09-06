package analysis

import "fmt"

type Correction interface {
	Correct(content string) string
}

func initCorrections() {
	MaterialCorrections[25] = &IdMaterial{}
}

var MaterialCorrections []Correction

type Material struct{}
type IdMaterial struct {}

func (material *IdMaterial) Correct(content string) string {
	fmt.Println(content)
	temp := []rune(content)
	position := 0
	for position < len(temp) {
		if temp[position] == 32 || temp[position] == 10 || temp[position] == 41 || temp[position] == 40 {
			break
		}
		position++
	}
	return string(temp[0:position])
}

