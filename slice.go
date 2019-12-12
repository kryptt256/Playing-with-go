package slices

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string
	numberArray := make([]int, 0, 3)

	for input != "X" {
		fmt.Print("Enter a number: ")
		fmt.Scan(&input)
		number, err := strconv.Atoi(input)
		if err == nil {
			if len(numberArray) >= 2 {
				numberArray = append(numberArray, number)
				sort.Ints(numberArray)
				fmt.Println("\nSorted numbers: ", numberArray)
			} else {
				numberArray = append(numberArray, number)
			}

		}
	}

}
