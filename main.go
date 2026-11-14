package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to the word counter!")

	fmt.Println("Enter file name: ")
	var fileName string
	fmt.Scanln(&fileName)

	fmt.Println("\nReading file...")
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

}
