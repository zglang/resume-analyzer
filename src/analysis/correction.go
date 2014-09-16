package analysis

import "fmt"

type Correction interface {
	Correct(content string) string
}

func initCorrections() {
	MaterialCorrections[16] = &IdentityMaterial{}
	MaterialCorrections[25] = &IdMaterial{}
}

var MaterialCorrections []Correction

type Material struct{}
type IdMaterial struct {}
type IdentityMaterial struct {}

func (material *IdMaterial) Correct(content string) string {

	fmt.Println("ID 处理",content)
	temp := []rune(content)
	fmt.Println(temp)
	position := 0
	for position < len(temp) {
		if temp[position] == 9 || temp[position] == 32 || temp[position] == 10 || temp[position] == 41 || temp[position] == 40 {
			break
		}
		position++
	}
	return string(temp[0:position])
}

func (material *IdentityMaterial) Correct(content string) string {

	fmt.Println("身份证 处理",content)
	temp := []rune(content)
	fmt.Println(temp)
	position := 0

	for position < len(temp) {
		if temp[position]>=97 && temp[position]<=122{
			position++
			break
		}
		if (temp[position] == 9 || temp[position] == 32 || temp[position] == 10) && position>10 {
			break
		}
		position++
	}
	return string(temp[0:position])
}

