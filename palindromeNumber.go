package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(0))
}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	if x < 10 && x > 0 {
		return true
	}

	var last, temp, a int

	a = x
	for x > 0 {
		last = x % 10
		temp = temp*10 + last
		x = x / 10
	}

	if a == temp {
		return true
	}

	return false
}
