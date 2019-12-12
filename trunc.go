package trunc

import "fmt"

func main() {
	var floatNumber float64
	var intNumber int64

	fmt.Print("Enter a floating point number: ")
	fmt.Scan(&floatNumber)
	intNumber = int64(floatNumber)
	fmt.Printf("The truncated version of %f is %d", floatNumber, intNumber)
}
