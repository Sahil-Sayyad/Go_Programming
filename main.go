package main

import (
	"fmt"
	"os"
)

func main(){

	file, err := os.Create("example.txt")
    
	if err != nil {
		fmt.Println("Errror", err)
		return
	}

	defer file.Close()

	content := "\nthis is file "

	_,err = file.WriteString(content)

	if err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println("content added successfully ")


}