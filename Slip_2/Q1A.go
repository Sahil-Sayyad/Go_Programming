/*
	A) Write a program in GO language to print Fibonacci series of n terms.
*/
package main

import "fmt"

func main(){

	var n int
	fmt.Println("Enter n number")
	fmt.Scan(&n)

	var a int = 0
	var b int = 1
    var fib int 
	fmt.Print(a)
	fmt.Print(" ",b)
	
	for i:=2;i<=n;i++ {

		fib = a+b
		a = b
		b = fib

		fmt.Print(" ", fib)

	}
}