package done

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	MAXSIZE = 4
	EXIT    = "X"
)

func ReadInts() []int {
	intSlice := make([]int, 0, 0)
	var input string

	for {
		fmt.Print("Enter a number or press 'X' to finish: ")
		fmt.Scan(&input)
		if strings.EqualFold(input, EXIT) {
			break
		}
		n, err := strconv.Atoi(input)

		if err == nil {
			intSlice = append(intSlice, n)
		} else {
			fmt.Printf("Error: %g, input again\n", err)
		}

	}
	return intSlice
}

func doPartition(numbers []int) map[string][]int {
	size := len(numbers)
	partitionSize := size / MAXSIZE

	arrayMap := map[string][]int{}
	i := 0
	j := partitionSize

	for _, ch := range "abc" {
		arrayMap[string(ch)] = numbers[i : partitionSize+i]
		i += partitionSize
		j += partitionSize
	}
	arrayMap["d"] = numbers[j-partitionSize:]

	return arrayMap
}

func sortingTask(slice []int, wg *sync.WaitGroup) {
	fmt.Println("Ordering sub-array ", slice)
	if !sort.IntsAreSorted(slice) {
		sort.Ints(slice)
	}
	wg.Done()
}

func mergeArrays(arr1 []int, arr2 []int, n1 int, n2 int) []int {
	arr3 := make([]int, n1+n2, n1+n2)
	var i, j, k int

	for i < n1 && j < n2 {
		if arr1[i] < arr2[j] {
			arr3[k] = arr1[i]
			k++
			i++
		} else {
			arr3[k] = arr2[j]
			k++
			j++
		}
	}

	for i < n1 {
		arr3[k] = arr1[i]
		k++
		i++
	}

	for j < n2 {
		arr3[k] = arr2[j]
		k++
		j++
	}

	return arr3
}

func merge(mapArray map[string][]int) []int {
	arr1, arr2, arr3, arr4 := mapArray["a"], mapArray["b"], mapArray["c"], mapArray["d"]

	half1 := mergeArrays(arr1, arr2, len(arr1), len(arr2))
	half2 := mergeArrays(arr3, arr4, len(arr3), len(arr4))

	return mergeArrays(half1, half2, len(half1), len(half2))
}

func main() {
	numbers := ReadInts()
	arrayMap := doPartition(numbers)
	var wg sync.WaitGroup

	for _, v := range arrayMap {
		if len(v) > 0 {
			go sortingTask(v, &wg)
			wg.Add(1)
		}
	}
	wg.Wait()

	fmt.Println("Ordered numbers: ", merge(arrayMap))
}
