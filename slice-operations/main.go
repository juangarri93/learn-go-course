package main

import (
	"fmt"
)

func main(){
	slice := []int{42, 43, 44}

	fmt.Printf("Initial Slice - %#v\n", slice)
	fmt.Println("-------------")

	// Appending values to a Slice
	slice = append(slice, 45, 46, 47, 99, 777)
	fmt.Printf("Slice after appending - %#v\n", slice)
	fmt.Println("-------------")

	// Slice a Slice
	fmt.Printf("Slice [inclusive:exclusive] - %#v\n", slice[0:4])
	fmt.Println("-------------")
	fmt.Printf("Slice [:exclusive] - %#v\n", slice[:7])
	fmt.Println("-------------")
	fmt.Printf("Slice [inclusive:] - %#v\n", slice[4:])
	fmt.Println("-------------")
	fmt.Printf("Slice [:] - %#v\n", slice[:])
	fmt.Println("-------------")

	// Delete from Slice
	slice = append(slice[:4], slice[5:]...)
	fmt.Printf("Slice after deleting \"46\" value - %#v\n", slice)

	fmt.Println("-------------")
	fmt.Println("-------------")

	// Make a Slide
	emptySlice := make([]int, 0, 10)

	fmt.Printf("Empty Slice - %#v\n", slice)
	fmt.Printf("Slice - len: %d - cap: %d\n", len(emptySlice), cap(emptySlice))
	
	fmt.Println("------------")

	emptySlice = append(emptySlice, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("Appended Slice - %#v\n", slice)
	fmt.Printf("Slice - len: %d - cap: %d\n", len(emptySlice), cap(emptySlice))
	fmt.Println("------------")

	emptySlice = append(emptySlice, 10, 11, 12, 13)
	fmt.Printf("Appended Slice - %#v\n", slice)
	fmt.Printf("Slice - len: %d - cap: %d\n", len(emptySlice), cap(emptySlice))
	fmt.Println("------------")
}