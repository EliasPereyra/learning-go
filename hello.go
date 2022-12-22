package main

import (
	"fmt"
);

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

}
