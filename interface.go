package main 

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width , length float64 
}

func (r Rectangle) Area() float64{
	return r.width* r.length
}


func main(){

	rect := Rectangle {
		width: 5,
		length: 3,
	}

	var shape Shape = rect

	if rectValue , ok := shape.(Rectangle); ok {

		fmt.Println("Interface value holds a Rectangle : ")
		fmt.Println("Width : ",rectValue.width)
		fmt.Println("height : ", rectValue.length)
		fmt.Println("Area : ", rectValue.Area())
	}else{
		fmt.Println("Interface value does not hold a Rectangle.")
	}

}