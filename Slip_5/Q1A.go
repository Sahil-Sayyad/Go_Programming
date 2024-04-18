/*
	A) Write a program in GO language program to create Text file 
*/

package main
import (
	"fmt"
	"os"
)

func main(){

	file , err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close();

	_, err = file.WriteString("Hello Sahil!")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("File created successfully!")
}