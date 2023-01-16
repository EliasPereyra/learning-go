package main

import (
	"fmt"
)

func main() {
	var language string = "Learning Go!"

	var copy string = language[:]

	const HTTP_GET_VERB, HTTP_PUT_VERB, HTTP_DELETE_VERB = "GET", "PUT", "DELETE"

	name, lastname, email := "Elias", "Pereyra", "elias@hotmail.com"
	var userInfo string = name + " " + lastname + " " + email

	fmt.Println("Hello world!")
	fmt.Println("You can build this file and compile it for running anywhere")
	fmt.Println("But you can specify the OS for which the file will be compiled to using the prefix 'GOOS' for Windows and 'GOARCH' for Linux")

	fmt.Println(HTTP_DELETE_VERB)
	fmt.Println("The user is: ", userInfo, "Shortname:", name[0:3])
	fmt.Println(language, len(language), "is the amount of bytes in memory")
	fmt.Println("A copy", copy)

	// If I know the amount of elements in an array, and the elements must be of the same type
	// Arrays cannot be resized
	var myarray = [3]string{"first", "second", "third"}

	// If I dunno what the amount is, Go does that check for me
	var yourarray = [...]string{"first", "last"}

	// I can copy an array
	newarray := myarray

	fmt.Println(yourarray, newarray)

	// A slice is a type of array, i.e., it's built on top of an array, the difference they have is that a slice is rezisable
	// An array is lower level than slice is, slice is comparable to an array of a high-level language like JS
	// It's not necessary to specify a length for a slice
	// var slice = []string{"one", "two"}
	myslice := make([]string, 3)
	myslice = append(myslice, "three", "four")

	// we can modify the slice's capacity so that we can extend it.
	// We can specify the capacity adding a third parameter to make():
	var some_names = make([]string, 0, 10)
	some_names = append(some_names, "elias", "pedro", "martin", "pablito")

	last_two_names := some_names[:2]
	first_two_names := some_names[2:]
	middle_names := some_names[1:3]

	fmt.Println(some_names)
	fmt.Println(last_two_names)
	fmt.Println(first_two_names)
	fmt.Println(middle_names)

	// a map does not need an amount of items
	agesMap := make(map[string]int)
	agesMap["elias"] = 25
	// or it can be written in this way
	ageMap := map[string]int{"elias": 25}
	fmt.Println(ageMap)
	delete(agesMap, "elias")

	// In go there are no whiles, there are only for loops
	var numbers = make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}

	// to iteare an array
	for i, num := range numbers {
		fmt.Printf("Number: %d - Position: %d\n", i, num)
	}
}
