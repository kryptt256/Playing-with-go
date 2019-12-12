package main1

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name, addr string
	fmt.Println("Enter a name: ")
	name = readInput()
	fmt.Println("Enter an address: ")
	addr = readInput()

	person := map[string]string{"name": name, "address": addr}
	data, _ := json.Marshal(person)
	fmt.Println(string(data))
}

func readInput() string {
	result := ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		result = scanner.Text()
	}
	return result
}
