/*
  B) Write a program in GO language to print file information. 
*/
package main

import (
	"fmt"
	"os"
	"log"
)

func main(){
	
	//specify the file path
	filePath := "example.txt"
   
	//open the file 
	file, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err)
	}

	//Print File information
	fmt.Println("File Name : ", file.Name());
	fmt.Println("Size (in bytes) : ", file.Size())
	fmt.Println("Permissions : ", file.Mode())
	fmt.Println("Last Modified : ", file.ModTime())
}