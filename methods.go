package main

import "fmt"

type rect struct{
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main(){
	r1 := rect{width: 10, height: 5}

	fmt.Println("area: ",r1.area())
	fmt.Println("perim: ", r1.perim())

	rp := &r1
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}


