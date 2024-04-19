/*
 A) Write a program in GO language to accept the book details such
	as BookID, Title, Author, Price. Read and display the details of
	‘n’ number of books
*/
package main

import "fmt"

type Book struct {
	BookID int 
	Title string 
	Author string 
	Price float64 

}

func main(){
	
	var n int;
	fmt.Println("enter n number of books ")
	fmt.Scanln(&n)
	book := make([]Book, n)

	for i:= 0; i < len(book); i++{
		fmt.Println("Enter book : ", i + 1 )
		fmt.Println("Enter Book id :")
		fmt.Scanln(&book[i].BookID)
		fmt.Println("Enter Title : ")
		fmt.Scanln(&book[i].Title)
		fmt.Println("Enter Author :")
		fmt.Scanln(&book[i].Author)
		fmt.Println("Enter Price : ")
		fmt.Scanln(&book[i].Price)
	}

	for _, b:= range book{
		fmt.Printf("\nBook Id : %d Book Title :%s  Book Author : %s Book Price : %.2f ", b.BookID,b.Title,b.Author,b.Price)

	}
}