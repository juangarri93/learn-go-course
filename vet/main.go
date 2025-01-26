package main

import (
	"strconv"
	"github.com/juangarri93/learn-go-course-dependency"
	"fmt"
	"math/rand"
	"time"
)

const vetName = "Puppies Vet"
var puppiesAmount = 3
const parralelBarkingEnabled = true
const rangeBarksEnabled = true

func init(){
	println("The vet is opening")
}

func main(){
	puppiesAmount++
	fmt.Println("In: " + vetName + " the " + strconv.Itoa(puppiesAmount) + " puppies will bark: " + puppy.Bark())
	fmt.Println("They can also super bark: " + puppy.SuperBark())

	if parralelBarkingEnabled {
		parralelBarking()
	}

	if rangeBarksEnabled {
		barkRange()
	}

	puppy := puppy.Puppy {
		Name:  "Pepe",
		Breed: "Border collie",
		Age: 4,
	}

	fmt.Printf("The puppy - name: %s - breed: %s - age: %d\n", puppy.Name, puppy.Breed, puppy.Age)
	fmt.Println(puppy)
}

func parralelBarking() {
	// Concurrency -> select a channel
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Random duration
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- rand.Int()
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- rand.Int()
	}()

	// A "select" statement chooses which of a set of possible send or receive operations will proceed.
	select {
		case v1 := <- ch1:
			fmt.Println("The bark came from channel 1: ", v1)
		case v2 := <- ch2:
			fmt.Println("The bark came from channel 2: ", v2)
	}
}

func barkRange() {
	// For range loop -> Data structures - slice
	xi := []string{"guf", "wof", "waf"}
	for i, v := range xi {
		fmt.Println("Ranging over barks", i, v)
	}
}