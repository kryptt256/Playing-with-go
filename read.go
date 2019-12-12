package main3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FIELDSIZE = 20

type name struct {
	fname [FIELDSIZE]byte
	lname [FIELDSIZE]byte
}

func main() {
	names := make([]name, 0)
	fmt.Print("Enter an existing file name: ")
	filename := readInput()
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sline := string(scanner.Text())
		fullName := strings.Split(sline, " ")
		name := name{}
		copy(name.fname[:], fullName[0])
		copy(name.lname[:], fullName[1])
		names = append(names, name)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	printNames(names)
}

func printNames(names []name) {
	for _, v := range names {
		s1 := string(v.fname[:FIELDSIZE])
		s2 := string(v.lname[:FIELDSIZE])
		fmt.Println(s1, s2)
	}
}

func readInput() string {
	result := ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		result = scanner.Text()
	}
	return result
}
