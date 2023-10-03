package mathformula

import (
	"fmt"
	"math"
)

// calculate area of a rectangle
func AreaOfRectangle(length float64, width float64) (float64, error) {
	if length < 0 || width < 0 {
		return 0, fmt.Errorf("input number cannot be less than zero(0)")
	}
	area_rec := length * width
	return area_rec, nil
}

// calculate the perimeter of  rectangle
func PerimeterOfRectangle(length float64, breadth float64) (float64, error) {
	if length < 0 && breadth < 0 {
		return 0, fmt.Errorf("input number cannot be less than zero(0)")
	}
	var peri = 2 * (length + breadth)
	return peri, nil
}

// calculate the area of a square
func AreaOfSquare(length float64) (float64, error) {
	if length < 0 {
		return 0, fmt.Errorf("input number cannot be less than zero(0)")
	}
	square_area := math.Pow(length, 2)
	return square_area, nil

}

// calculate the perimeter of a square
func PerimeterOfSquare(length float64) (float64, error) {
	if length < 0 {
		return 0, fmt.Errorf("input number cannot be less than zero(0)")
	}
	perimeter_sq := 4 * (length)
	return perimeter_sq, nil
}

// calculate the area of a circle
func AreaOfCircle(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("input number cannot be less than zero(0)")
	}
	circle_area := math.Pi * math.Pow(radius, 2)
	return circle_area, nil
}

// calculating the area of triangle
func AreaOftriangle(base float64, height float64) (float64, error) {
	if base < 0 || height < 0 {
		return 0, fmt.Errorf("input number cannot be less than zero(0)")
	}
	area_tri := (base * height) / 2
	return area_tri, nil
}

//sending momo

func CurrentBalance(availablebal float64, payment float64) (float64, error) {
	if payment > availablebal {
		return 0, fmt.Errorf("insufficient balance")

	}
	
	current_bal := availablebal - payment

	return current_bal, nil

}
