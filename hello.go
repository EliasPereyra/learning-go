package main

import "fmt";

func main() {
	var language string = "Learning Go!"
	var copy string = language[:]

	const HTTP_GET_VERB, HTTP_PUT_VERB, HTTP_DELETE_VERB = "GET", "PUT", "DELETE"

	name, lastname, email := "Elias", "Pereyra", "elias@hotmail.com"

	fmt.Println("Hello world!")
	fmt.Println("You can build this file and compile it for running anywhere")
	fmt.Println("But you can specify the OS for which the file will be compiled to using the prefix 'GOOS' for Windows and 'GOARCH' for Linux")

	fmt.Println(HTTP_DELETE_VERB)
	fmt.Println("The user is: ", name,lastname, "Shortname:", name[0:3], email)
	fmt.Println(language, len(language), "is the amount of bytes in memory")
	fmt.Println("A copy", copy)
}
