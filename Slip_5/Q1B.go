/*
B) Write a program in GO language to accept n records of employee information 
   (eno,ename,salary) and display records of employees having minimum salary.

*/
package main 
import "fmt"

type Employee struct {
	eno int 
	ename string 
	salary int 
}

func main(){

     var n int;
	 fmt.Println("Enter How many employees : ")
	 fmt.Scan(&n)

	 employees := make([]Employee, n)

	 for i := 0 ; i < n; i++ {

		fmt.Println("Enter Employee " , i + 1);
		fmt.Println("Enter eno : ")
		fmt.Scanln(&employees[i].eno)
		fmt.Println("Enter ename : ")
		fmt.Scanln(&employees[i].ename)
		fmt.Println("Enter salary : ")
		fmt.Scanln(&employees[i].salary)

	 }

	 maxSalary := employees[0].salary

	 for _, emp := range employees{
		if emp.salary > maxSalary {
			maxSalary = emp.salary;
		}
	 }

	 for _, emp := range employees{
		if emp.salary == maxSalary{
			fmt.Println("\nEmployee NO :  Employee Name : Employee Salary : ", emp.eno, emp.ename,emp.salary)
		}
	 }
}