package findian

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string
	fmt.Print("Enter a string: ")
	str = strings.ToLower(readData())

	if isStringValid(str) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func readData() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use this to keep reading
	// get the line after user press <enter>
	return scanner.Text()
}

func isStringValid(s string) bool {
	return strings.HasPrefix(s, "i") && strings.ContainsRune(s, 'a') && strings.HasSuffix(s, "n")
}
