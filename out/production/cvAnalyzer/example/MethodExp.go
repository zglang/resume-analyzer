package example

import "math"
import (
	"fmt"
)

const (
	WHITE=iota
	BLANK
	BLUE
	RED
)

type Color byte

type Shape struct {
	color Color
}

type Star struct {
	angle int
	shape Shape
}

type Rectangle struct {
	width,height float64
	Shape
}

type Circle struct {
	radius float64
	Shape
}

func (s Shape) printColor(){
	fmt.Printf("color=%d\n",s.color)
}

func (r Rectangle) area() float64{
	return r.height*r.width
}

func (c Circle) area() float64{
	return c.radius * c.radius * math.Pi
}

func (r *Rectangle) setWidth(width float64){
	r.width=width
	fmt.Printf("*func--width=%f,height=%f,Rectangle.area()=%f\n",r.width,r.height,r.area())
}

func MethodTest4(){
	c:=Circle{100,Shape{3}}
	s:=Star{1,Shape{RED}}
	s.shape.printColor();
	fmt.Printf("color=%d,radius=%f,Circle.area()=%f\n",c.color, c.radius,c.area())
	c.printColor()
}

func MethodTest3(){
	r:=[]Rectangle{{100,100,Shape{RED}},{200,200,Shape{RED}}}
	for index,value:=range r{
		value.setWidth(500)
		fmt.Printf("[%d] width=%f,height=%f,Rectangle.area()=%f\n",index,value.width,value.height,value.area())
	}
}

func MethodTest2(){
	r:=Rectangle{100,100,Shape{RED}}
	fmt.Printf("width=%f,height=%f,Rectangle.area()=%f\n",r.width,r.height,r.area())
	r.setWidth(200)
	fmt.Printf("width=%f,height=%f,Rectangle.area()=%f\n",r.width,r.height,r.area())
}


func MethodTest1(){
	r:=Rectangle{100.6,22.7,Shape{RED}}
	c:=Circle{22.9,Shape{RED}}
	fmt.Printf("Rectangle.area()=%f\n",r.area())
	fmt.Printf("Circle.area()=%f\n",c.area())
}
