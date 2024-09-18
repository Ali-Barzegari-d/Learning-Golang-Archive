package main

import "fmt"

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome to World,", firstName)
}

func getNames() (string, string) {
	return "Ali", "Barzeegari"
}
