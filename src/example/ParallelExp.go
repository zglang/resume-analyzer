package example

import (
	"fmt"
	"runtime"
)

func hello(a,b,c int){
	fmt.Printf("++++++++++++start+++++++++++++++")
	for i:=0;i<5;i++{
		runtime.Gosched()
		fmt.Printf("[%d] a,b,c=%d,%d,%d\n",i,a,b,c)
	}
	fmt.Printf("++++++++++++end+++++++++++++++")
}

func ParallelSelect1(){

	ch:=make(chan int)
	qch:=make(chan int)
	go func(){
		for i:=0;i<10;i++{
			fmt.Println(<-ch)
		}
		qch <- 0
	}()
	x,y:=1,1
	for{
		select{
			case ch <- x:
				x,y=y,x+y
			case <-qch:
				fmt.Println("quit")
				return
		default:
			runtime.Gosched()
			fmt.Println("----")
		}
	}

}

func ParallelChannel1(){
	a:=[]int{1,2,3,4,5,6,7,8,9}
	ch := make(chan int)
	go sum(a[:],ch)
	go sum(a[5:],ch)
	x,y:=<-ch,<-ch
	fmt.Printf("x=%d,y=%d\n",x,y)

}

func ParallelChannel2(){
	ch:=make(chan int,2)
	ch <- 5
	ch <- 6
	close(ch)
	for c:=range ch{
		fmt.Printf("ch=%d\n",c)
	}
	//close(ch)
}

func sum(a []int, c chan int){
	sum:=0;
	for _,v:=range a{
		sum+=v
	}
	c <- sum
}


func ParallelTest1(){
	fmt.Printf("++++++++++++ParallelTest1 start++++++++++++\n")
	go hello(5,4,6)
	hello(1,2,3)
	fmt.Printf("++++++++++++ParallelTest1 end++++++++++++\n")
}

