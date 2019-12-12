package done

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	COMMANDS  = "newanimal query"
	ANIMALS   = "cow bird snake"
	CREATEMSG = "Created it!"
	INFO      = "Valid info: eat, move, speak"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AAnimal struct {
	name, food, locomotion, noise string
}

func (a AAnimal) getName() string {
	return a.name
}

type Cow AAnimal
type Bird AAnimal
type Snake AAnimal

// Implementation for Cow

func (cow Cow) Eat() {
	fmt.Println(cow.food)
}

func (cow Cow) Move() {
	fmt.Println(cow.locomotion)
}

func (cow Cow) Speak() {
	fmt.Println(cow.noise)
}

//===========================

// Implementation for Bird
func (bird Bird) Eat() {
	fmt.Println(bird.food)
}

func (bird Bird) Move() {
	fmt.Println(bird.locomotion)
}

func (bird Bird) Speak() {
	fmt.Println(bird.noise)
}

//==============================

// Implementation for Snake
func (snake Snake) Eat() {
	fmt.Println(snake.food)
}

func (snake Snake) Move() {
	fmt.Println(snake.locomotion)
}

func (snake Snake) Speak() {
	fmt.Println(snake.noise)
}

func CreateAnimal(name string, animalType string) Animal {
	var a Animal
	switch animalType {
	case "cow":
		a = Cow{name, "grass", "walk", "moo"}
	case "bird":
		a = Bird{name, "worms", "fly", "peep"}
	default:
		a = Snake{name, "mice", "slither", "hsss"}
	}

	return a
}

func findByName(name string, animals []Animal) Animal {
	var result Animal
	for _, v := range animals {
		switch sh := v.(type) {
		case Cow:
			if AAnimal(sh).getName() == name {
				result = v
				break
			}
		case Bird:
			if AAnimal(sh).getName() == name {
				result = v
				break
			}
		case Snake:
			if AAnimal(sh).getName() == name {
				result = v
				break
			}
		}
	}

	return result
}

func makeRequest() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var err error
	var fieldList []string

	if scanner.Scan() {
		str := strings.ToLower(scanner.Text())
		fieldList = strings.Fields(str)
		err = checkRequest(fieldList)
	}

	return fieldList, err
}

func checkRequest(fieldList []string) error {
	if len(fieldList) < 3 {
		return errors.New("Each command must contains at least 3 strings")
	}

	if !strings.Contains(COMMANDS, fieldList[0]) {
		return errors.New("Invalid command")
	}

	if fieldList[0] == "newanimal" && !strings.Contains(ANIMALS, fieldList[2]) {
		return errors.New("Invalid animal type\n" + "Valid animals are: " + ANIMALS)
	}

	if fieldList[0] == "query" && !strings.Contains(INFO, fieldList[2]) {
		return errors.New("Invalid requested info\n" + INFO)
	}

	return nil
}

func processQuery(animal Animal, action string) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	default:
		animal.Speak()
	}
}

func main() {
	animals := make([]Animal, 0, 0)

	for {
		fmt.Print("> ")
		fieldList, err := makeRequest()

		if err == nil {
			name := fieldList[1]
			switch fieldList[0] {
			case "newanimal":
				a := findByName(name, animals)
				if a != nil {
					fmt.Println("Error: An animal with name ", name, " already exists.")
				} else {
					animals = append(animals, CreateAnimal(name, fieldList[2]))
				}

			default:
				a := findByName(name, animals)
				if a != nil {
					processQuery(a, fieldList[2])
				} else {
					fmt.Println("This animal doesn't exist")
				}

			}
		} else {
			fmt.Println("Wrong request: ", err)
		}

	}

}
