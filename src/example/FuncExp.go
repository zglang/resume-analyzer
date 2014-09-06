package example

import "fmt"

func MultResult1(A,B int) (add int,mult int){
	add = A+B;
	mult = A*B;
	return;
}

func MultResult2(A int,B int) (int,int){
	add := A+B;
	mult := A*B;
	return add,mult;
}

func DynamicParam(args ...int){
	for _,v := range args{
		fmt.Println(v)
	}
}

type testInt func(int,int) (int,int)

func TestFunc(aa []int,t testInt){
	for _,v:= range aa{
		a,b:=t(v,100)
		fmt.Printf("a=%d,b=%d\n",a,b)
	}
}





