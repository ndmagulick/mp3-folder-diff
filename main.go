package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")
	readDirectory()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readDirectory() {
	directoryEntries, err := os.ReadDir("")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range directoryEntries {
		if !entry.IsDir() {
			fmt.Println(" ", entry.Name())
		}
	}
}
