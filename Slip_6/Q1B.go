/*
  B) Write a program in GO language to copy all elements of one array
     into another using a method.
*/

package main

import "fmt"

func copyArray(source[]int, destination[]int){

	for i:=0; i < len(source); i++{
		destination[i] = source[i]
	}
	
}
func main(){

	sourceArray := []int {1,2,3,4,5}

	destinationArray := make([]int, len(sourceArray))
	
	fmt.Println("Source Array : ", sourceArray)

	copyArray(sourceArray,destinationArray)

	fmt.Println("Destination array after copying ", destinationArray)
}