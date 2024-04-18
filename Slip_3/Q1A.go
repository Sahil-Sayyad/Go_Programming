/*
	A) Write a program in the GO language using function to check whether 
	   accepts number is palindrome or not.
*/

package main 

import "fmt"

func isPalindrome(n int){

	var temp int
    
	temp = n
 
	var check int
	for  temp != 0{
      
		var digit = temp % 10

		check = check * 10 + digit

		temp /=10
	}

	if(check==n){
		fmt.Println(n," is palindrome.")
	}else{
		fmt.Println(n,"is not palindrome.")
	}
}
func main(){

	var n int
	fmt.Println("Enter number to check is palindrome or not")
	fmt.Scan(&n)

	isPalindrome(n)
}