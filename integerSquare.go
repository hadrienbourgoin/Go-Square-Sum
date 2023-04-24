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
		var casenbr int
		for i > 0 {
			i--
			casenbr++
			fmt.Println("Test case number: ", casenbr, ". Choose the number of integers you want: ")
			fmt.Scanln(&s)
			if b2, i2 := controluserinput(s); b2 {
				fmt.Println("You'll have to enter ", i2, " integers: ")
				var sum int
				for i2 > 0 {
					i2--
					fmt.Println("Enter an integer: ")
					fmt.Scanln(&s)
					if b3, i3 := controluserinput(s); b3 {
						sum = sum + i3*i3
					} else {
						fmt.Println("It wasn't an integer.")
					}
				}
				fmt.Println("The sum of square root of all your integers is: ", sum)
			} else {
				fmt.Println("It wasn't an integer.")
			}
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
