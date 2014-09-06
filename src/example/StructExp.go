package example

import (
	"fmt"
)

type Person struct{
	name string
	age int
}

type Stills []int

type Months map[string]int

type Student struct {
	class int
	Person
	Months
	Stills
	int
}

func init(){
	fmt.Printf("+++++++++++++++++++++++++\n")
}

func TestStruct4(){
	var P2 Student
	P2.Stills=Stills{1,2,3,5,6}
	for _,v:=range P2.Stills{
		fmt.Println(v)
	}

}



func TestStruct3(){
	P2:=Student{1,Person{"laowei",33},Months{"a":1,"b":2},Stills{1,2,3},99}
	fmt.Printf("P2.name=%s,age=%d,class=%d\n",P2.name,P2.age,P2.class)
	fmt.Printf("P2.name=%s,age=%d,class=%d\n",P2.Person.name,P2.Person.age,P2.class)
	for k,v:=range P2.Months{
		fmt.Printf("k=%s,v=%d\n",k,v)
	}
	P2.Months = Months{"c":222}
	for k,v:=range P2.Months{
		fmt.Printf("k=%s,v=%d\n",k,v)
	}
	for _,v:=range P2.Stills{
		fmt.Println(v)
	}
	P2.Stills = Stills{33,22,11}
	for _,v:=range P2.Stills{
		fmt.Println(v)
	}
	fmt.Printf("P2.int=%d",P2.int)
}

//func TestStruct2(){
//	P2:=Student{1,Person{"laowei",33}}
//	fmt.Printf("P2.name=%s,age=%d,class=%d\n",P2.name,P2.age,P2.class)
//	fmt.Printf("P2.name=%s,age=%d,class=%d\n",P2.Person.name,P2.Person.age,P2.class)
//}

func TestStruct1(){
	var P Person
	P.age=10
	P.name="laowei"

	P2:= Person{"xiaohong",20}
	fmt.Printf("P2.name=%s\n",P2.name)

	P3:= Person{age:11,name:"liudehua"}
	fmt.Printf("P2.name=%s\n",P3.name)
}

