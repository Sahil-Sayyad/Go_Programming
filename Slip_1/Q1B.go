/*
	B) Write a program in GO language to accept n student details like roll_no,
	   stud_name, mark1,mark2, mark3. Calculate the total and average of marks. 
	   using structure.
*/
package main

import  "fmt"

type Student struct{
	roll_no int
	stud_name string
	mark1 int
	mark2 int
	mark3 int
}


func main(){

    var student Student
	
	fmt.Println("Enter Student Roll Number : ")
	fmt.Scan(&student.stud_name);
	fmt.Println("Enter Student Name :")
	fmt.Scan(&student.stud_name)
	fmt.Println("Enter Mark1 : ")
	fmt.Scan(&student.mark1);
	fmt.Println("Enter Mark2 : ")
	fmt.Scan(&student.mark2)
	fmt.Println("Enter Mark3 : ")
	fmt.Scan(&student.mark3)

    var result = (student.mark1 + student.mark2 + student.mark3) / 3;


	fmt.Println("Average of Marks : ", result)
}


