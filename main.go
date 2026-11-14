package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to the word counter!")

	fmt.Println("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	fmt.Println("\nReading file...")
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println("File name: " + fileName)
	fmt.Printf("File size: %d", len(data))
	fmt.Printf("File content: %s", data)

}
