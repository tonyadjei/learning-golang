package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	// get the reference of a variable: use &
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)

}

	// indicate that a type is a pointer, use * in front of the type
func changeUsingPointer(s *string) {
	// this will print out the actually memory address/location
	log.Println("s is set to", s)
	newValue := "Red"
	// access the content/value stored at a memory location, and modify it, put an * in front of the pointer
	*s = newValue
}
