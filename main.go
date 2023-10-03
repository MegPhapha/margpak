package main

import (
	"fmt"
	"margpak/mathformula"
	//"margpak/mathformula"
	//"margpak/utils"
)

func main() {
	// result, err := utils.CheckNumberAgainstExpected(4, "Noah")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// output := utils.NumberOfCharacters("Meg")
	// fmt.Println("Number of characters is:", output)

	// // calculate area of a rectangle
	// areaOfRectangle, err := mathformula.AreaOfRectangle(-2, 5)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("The area of the rectangle is:", areaOfRectangle)

	// // calculate the perimeter of  rectangle
	// perimeterOfRectangle, err := mathformula.PerimeterOfRectangle(3, 5)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("The perimeter of the rectangle is:", perimeterOfRectangle)

	// //calculate the area of a square
	// areaOfSquare, err := mathformula.AreaOfSquare(9)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("The area of the rectangle is:", areaOfSquare)

	// // calculate the perimeter of a square
	// perimeterOfSquare, err := mathformula.PerimeterOfSquare(3.5)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("The perimeter of the square:", perimeterOfSquare)

	// // calculate the area of a circle
	// areaOfCircle, err := mathformula.AreaOfCircle(6.4)
	// c
	// 	return
	// }
	// fmt.Println("The area of the circle is:", areaOfCircle)

	// //parse name to return person's name
	// fullname := utils.ParseName("Noah", "Armdress")
	// fmt.Println("Parsed name:", fullname)

var choice int
fmt.Println("Polygon Area Calculator")
fmt.Println("1. Rectangle")
fmt.Println("2. Triangle")
fmt.Println("3. Square")
fmt.Print("Enter your choice:")

fmt.Scanln(&choice)


	var area float64

	if choice == 1 {
		// Calculate the area of a rectangle
		var length, width float64
		fmt.Print("Enter the length of the rectangle: ")
		fmt.Scanln(&length)
		fmt.Print("Enter the width of the rectangle: ")
		fmt.Scanln(&width)
		area_rec, err := mathformula.AreaOfRectangle(length, width)
		if err != nil {
			fmt.Println(err)
			return
		}
		area = area_rec
		
		} else if choice == 2 {
		// Calculate the area of a triangle
		var base, height float64
		fmt.Print("Enter the base of the triangle: ")
		fmt.Scanln(&base)
		fmt.Print("Enter the height of the triangle: ")
		fmt.Scanln(&height)
		area_tri, err := mathformula.AreaOftriangle(base, height)
		if err != nil {
			fmt.Println(err)
	         return
		}
		area = area_tri

	} else if choice == 3 {
		// Calculate the area of a square
		var length float64
		fmt.Print("Enter the side length of the square: ")
		fmt.Scanln(&length)
		area_square, err:= mathformula.AreaOfSquare(length)
		if err != nil {
			fmt.Println(err)
			return
		}

		area = area_square

	} else {
		fmt.Println("Invalid choice")
		return
	}

	//fmt.Printf("The area of the polygon is %.2f\n", area)
	fmt.Println("The area of the polygon is", area)
}
