package main

import "fmt"

func vals() (int,int){
	return 3,7
}

func vals3() (int,int,string){
	return 1,2,"three"
}

func main(){
	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)
	
	_,c := vals()
	fmt.Println(c)

	x,y,z := vals3()

	fmt.Println(x,y,z)
}




