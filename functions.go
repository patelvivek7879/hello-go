package main

import "fmt"


func plus(a int,b int) int {
	return a+b
}

func plusPlus(a,b,c int) int{
	return a + b + c
}

func strFun(a string, b string) string {
	return a + b
}

func main(){
	res := plus(1,2)
	fmt.Println("1+2 =", res);

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3 =", res)

	res1 := strFun("vivek","patel");
	fmt.Println("vivek+patel = ", res1)
}



