package done

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	ANIMALS = "Valid animals: cow, bird, snake"
	INFO    = "Valid info: eat, move, speak"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animals := map[string]Animal{"cow": Animal{"grass", "walk", "moo"}, "bird": Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"}}

	for {
		fmt.Print("> ")
		name, action, err := makeRequest()

		if err == nil {
			animal := animals[name]
			answerRequest(animal, action)
		} else {
			fmt.Println("Wrong request: ", err)
		}

	}
}

func makeRequest() (string, string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var animalName, animalAction string
	var err error

	if scanner.Scan() {
		str := scanner.Text()
		animalName, animalAction, err = checkRequest(str)
	}

	return animalName, animalAction, err
}

func answerRequest(animal Animal, action string) {
	switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
	}
}

func checkRequest(requestStr string) (string, string, error) {
	requestStr = strings.ToLower(requestStr)
	line := strings.Split(requestStr, " ")
	if len(line) < 2 {
		return "", "", errors.New("The request must contains a space between animal and action")
	}
	name := line[0]
	info := line[1]
	if !strings.Contains(ANIMALS, name) || !strings.Contains(INFO, info) {
		errorMsg := strings.Join([]string{"Invalid animal or requested info", ANIMALS, INFO}, "\n")
		return "", "", errors.New(errorMsg)
	}

	return name, info, nil
}
