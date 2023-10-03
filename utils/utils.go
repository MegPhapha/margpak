
package utils

import "fmt"

func CheckNumberAgainstExpected(expectedLength int, v string) (bool, error) {
	if expectedLength == len(v) {
		return true, nil
	} else {
		return false, fmt.Errorf("input length does not match the expected length")
	}
}

func NumberOfCharacters(v string) int {

	return len(v)
}

// parse name in function and return the person's name
	// func ParseName(a, b string) string {
	// 	fullname := a + " " + b
	// 	return fullname
//}

func ParseName(a, b string) string {
	if a == "" {
		return b
	} else if b == "" {
		return a
	}
	fullname := a + " " + b
	return fullname
}




