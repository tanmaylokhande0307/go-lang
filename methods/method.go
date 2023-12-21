package main

import "fmt"

type rect struct{
	width , height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perimeter() int {
	return 2*r.height + 2*r.width
}

func main(){

	r := rect{height: 5,width: 10}

	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perimeter())

	fmt.Println(r)



}