package main1

import (
	"fmt"
)

func main() {
	var number int
	numberList := make([]int, 0, 10)

	for len(numberList) < 10 {
		fmt.Print("Enter a number: ")
		_, err := fmt.Scanf("%d", &number)
		if err == nil {
			numberList = append(numberList, number)
		}
	}
	fmt.Println("You entered the numbers: ", numberList)
	BubbleSort(numberList)
	fmt.Println("Sorted list: ", numberList)
}

func BubbleSort(numbers []int) {
	size := len(numbers)
	for {
		swapped := false
		for i := 1; i < size; i++ {
			if numbers[i-1] > numbers[i] {
				Swap(numbers, i-1)
				swapped = true
			}
		}
		size--
		if !swapped {
			break
		}
	}

}

func Swap(numbers []int, i int) {
	aux := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = aux
}
