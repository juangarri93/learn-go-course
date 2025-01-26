package main

import(
	"github.com/juangarri93/learn-go-course-dependency"
	"fmt"
)

func main(){
	puppy := puppy.Puppy {
		Name:  "Pepe",
		Breed: "Border collie",
		Age: 4,
	}

	fmt.Printf("The puppy - name: %s - breed: %s - age: %d\n", puppy.Name, puppy.Breed, puppy.Age)
	fmt.Println(puppy)

	anonymousStruct := struct {
		name string
		breed string 
		age int
	}{
		name:  "PepeClone",
		breed: "Border collie",
		age: 1,
	}

	fmt.Printf("The puppy - name: %s - breed: %s - age: %d\n", anonymousStruct.name, anonymousStruct.breed, anonymousStruct.age)
	fmt.Println(anonymousStruct)

	fmt.Println("--------------------------")
	fmt.Println("--------------------------")

	fmt.Println("Structure composition example")
	car := Car{}
	car.Start() // Calling the Start() method from the embedded Engine struct
}


type Engine struct {
	// Engine fields
}

// Engine method
func (e *Engine) Start() {
	fmt.Println("Engine started")
}

type Car struct {
	Engine // Embedding the Engine struct
	// Car-specific fields
}