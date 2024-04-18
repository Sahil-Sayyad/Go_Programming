/*
	A) Write a program in GO language to print a recursive sum of digits
	   of a given number.
*/

package main

import "fmt"

func recursiveDigitSum(n int) int {
    if n < 10 {
        return n
    }
    // Recursive step: Sum of digits of the number
    return n%10 + recursiveDigitSum(n/10)
}

func main() {
    var num int
    fmt.Println("Enter a number:")
    fmt.Scanln(&num)

    result := recursiveDigitSum(num)
    fmt.Printf("The recursive sum of digits of %d is: %d\n", num, result)
}
