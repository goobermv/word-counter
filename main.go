package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	frequency := make(map[string]int)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		word = strings.Trim(word, ".,!?;:\"'()[]{}")

		if word != "" {
			frequency[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
