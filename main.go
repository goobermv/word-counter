package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

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

	var ss []kv
	for k, v := range frequency {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	fmt.Println("\nTop 10 most frequent words:")

	count := 1
	for _, kv := range ss {
		if count > 10 {
			break
		}
		fmt.Printf("%d. %s - %d\n", count, kv.Key, kv.Value)
		count++
	}
}
