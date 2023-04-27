package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Println("Choose the number of test cases, it has to be an integer 1<=i<=100: ")
	fmt.Scanln(&s)
	if b, i := controluserinput(s); b {
		if i >= 1 && i <= 100 {
			looptestcase(i)
		} else {
			fmt.Println("It has to be between 1 and 100.")
		}
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
		fmt.Println("Choose the number of integers you want 1<=i<=100: ")
		fmt.Scanln(&s)
		if b, i := controluserinput(s); b {
			if i >= 1 && i <= 100 {
				n--
				fmt.Println("You'll have to enter ", i, " integers: ")
				loopsquare(i, 0)
			} else {
				fmt.Println("It has to be between 1 and 100.")
			}
		} else {
			fmt.Println("It wasn't an integer.")
		}
		looptestcase(n)
	}
}

func loopsquare(n int, sum int) {
	var s string
	if n > 0 {
		fmt.Println("Enter an integer -100<=i<=100: ")
		fmt.Scanln(&s)
		if b, i := controluserinput(s); b {
			if i >= -100 && i <= 100 {
				n--
				if i > 0 {
					sum = sum + i*i
				}
			} else {
				fmt.Println("It has to be between -100 and 100.")
			}
		} else {
			fmt.Println("It wasn't an integer.")
		}
		loopsquare(n, sum)
	} else {
		fmt.Println("The sum of square root of all your integers excluding negatives is: ", sum)
	}
}
