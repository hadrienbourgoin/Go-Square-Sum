package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Println("Choose the number of test cases, it has to be an integer: ")
	fmt.Scanln(&s)
	if b, i := controluserinput(s); b {
		looptestcase(i)
	} else {
		fmt.Println("It wasn't an integer.")
	}
}

func controluserinput(uinput string) (bool, int) {
	if i, err := strconv.Atoi(uinput); err == nil {
		return true, i
	} else {
		return false, i
	}
}

func looptestcase(n int) {
	var s string
	if n > 0 {
		fmt.Println("Choose the number of integers you want: ")
		fmt.Scanln(&s)
		if b, i := controluserinput(s); b {
			n--
			fmt.Println("You'll have to enter ", i, " integers: ")
			loopsquare(i, 0)
		} else {
			fmt.Println("It wasn't an integer.")
		}
		looptestcase(n)
	}
}

func loopsquare(n int, sum int) {
	var s string
	if n > 0 {
		fmt.Println("Enter an integer: ")
		fmt.Scanln(&s)
		if b, i := controluserinput(s); b {
			n--
			sum = sum + i*i
		} else {
			fmt.Println("It wasn't an integer.")
		}
		loopsquare(n, sum)
	} else {
		fmt.Println("The sum of square root of all your integers is: ", sum)
	}
}
