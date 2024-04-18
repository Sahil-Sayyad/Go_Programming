/*
B) Write a program in GO language to create structure student. 
   Write a method show() whose receiver is a pointer of struct student.
*/

package main 

import "fmt"

type Student struct {
	Name string 
	Age int 
	Grade string
}

func (s *Student) show(){
	fmt.Println("Name :", s.Name)
	fmt.Println("Age : ", s.Age)
	fmt.Println("Grade : ", s.Grade)
}

func main(){
	
	student := Student{
		Name : "Sahil Sayyad",
		Age : 22,
		Grade : "B",
	}

	student.show()
}