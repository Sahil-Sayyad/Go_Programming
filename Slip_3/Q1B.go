/*
B) Write a Program in GO language to accept n records of employee information 
(eno,ename,salary) and display record of employees having maximum salary.
*/

package main

import "fmt"

type Employee struct{
	eno int
	ename string
	salary float64
}

func main(){
   
	var n int 

	fmt.Println("How many employee : ")
	fmt.Scan(&n)

	employee := make([]Employee, n)


	for i := 0 ; i < n; i++ {

		fmt.Println("Enter eno :")
		fmt.Scan(&employee[i].eno)
		fmt.Println("Enter ename : ")
		fmt.Scan(&employee[i].ename)
		fmt.Println("Enter salary")
		fmt.Scan(&employee[i].salary)
	}

	maxSalary := employee[0].salary

	for _, emp := range employee{
		if emp.salary > maxSalary{
			maxSalary = emp.salary
		}
	}

	fmt.Println("\nEmployees with maximum Salary:")

	for _, emp := range employee{
		if emp.salary == maxSalary{
			fmt.Printf("Employee Number: %d, Employee Name : %s, Salary: %.2f\n", emp.eno,emp.ename,emp.salary)
		}
	}
}