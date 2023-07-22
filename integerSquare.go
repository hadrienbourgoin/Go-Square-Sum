package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
)

func main() {
	var nbrtestcase int = 0
	// fmt.Println("Choose the number of test cases, it has to be an integer 1<=i<=100:")
	fmt.Scanln(&nbrtestcase)
	var resultats []int
	if nbrtestcase >= 1 && nbrtestcase <= 100 {
		resultats = looptestcase(nbrtestcase, resultats)
	} else {
		fmt.Println("nbrtestcase not between 1 and 100")
	}
	displayresults(resultats, len(resultats))
}

func looptestcase(nbrloop int, arrayofresults []int) []int {
	if nbrloop > 0 {
		// fmt.Println("test case number: ", nbrloop, ". Choose the number of integers you want 0<i<=100:")
		var nbrintegers int = 0
		fmt.Scanln(&nbrintegers)
		if nbrintegers > 0 && nbrintegers <= 100 {
			// fmt.Println("You'll have to enter ", nbrintegers, " integers, space separated:")
			reader := bufio.NewReader(os.Stdin)
			userintegersstring, _ := reader.ReadString('\n')
			userintegersstring = strings.ReplaceAll(userintegersstring, "\r", "")
			userintegersstring = strings.ReplaceAll(userintegersstring, "\n", "")
			arrayofstring := strings.Split(userintegersstring, " ")
			if nbrintegers > len(arrayofstring) {
				nbrintegers = len(arrayofstring)
			}
			var arrayofint []int
			arrayofint = arraystringtoint(nbrintegers, arrayofstring, arrayofint)
			var result int = 0
			arrayofresults = append(arrayofresults, sumofintegers(arrayofint, nbrintegers, result))
			return looptestcase(nbrloop-1, arrayofresults)
		} else {
			fmt.Println("nbrintegers not between 0 and 100")
		}
	}
	return arrayofresults
}

func arraystringtoint(nbrloop int, arrayofstring []string, arrayofint []int) []int {
	if nbrloop > 0 {
		index := len(arrayofstring)-nbrloop
		if arrayofstring[index] != "" {
			intVal, err := strconv.Atoi(arrayofstring[index])
			if err != nil {
				fmt.Printf("Error converting string '%s' to integer: %v\n", arrayofstring[nbrloop], err)
				arrayofint = append(arrayofint, 0)
			} else {
				arrayofint = append(arrayofint, intVal)
			}
		} else {
			arrayofint = append(arrayofint, 0)
		}
		return arraystringtoint(nbrloop-1, arrayofstring, arrayofint)
	}
	return arrayofint
}

func sumofintegers(arrayofintegers []int, nbrloop int, result int) int {
	if nbrloop > 0 {
		index := len(arrayofintegers)-nbrloop
		if arrayofintegers[index] > 0 {
			result += arrayofintegers[index] * arrayofintegers[index]
		}
		return sumofintegers(arrayofintegers, nbrloop-1, result)
	}
	return result
}
func displayresults(arrayofresults []int, nbrloop int) {
	if nbrloop > 0 {
		index := len(arrayofresults)-nbrloop
		fmt.Println(arrayofresults[index])
		displayresults(arrayofresults, nbrloop-1)
	}
 }
