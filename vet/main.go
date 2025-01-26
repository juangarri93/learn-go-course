package main

import (
	"strconv"
	"github.com/juangarri93/learn-go-course-dependency"
)

const vetName = "Puppies Vet"
var puppiesAmount = 3

func main(){
	puppiesAmount++
	println("In: " + vetName + " the " + strconv.Itoa(puppiesAmount) + " puppies will bark:" + puppy.Bark())
	println("Or even super bark:" + puppy.SuperBark())
}

