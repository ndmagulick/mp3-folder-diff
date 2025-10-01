package main

import (
	"fmt"
	"os"
)

func main() {
	const sourceDirectory string = ""
	const destinationDirectory string = ""

	var sourceDirectoryEntries []string = readDirectory(sourceDirectory)
	printEntries(sourceDirectoryEntries)

	var destinationDirectoryEntries []string = readDirectory(destinationDirectory)
	printEntries(destinationDirectoryEntries)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readDirectory(sourceDirectory string) []string {
	directoryEntries, err := os.ReadDir(sourceDirectory)
	check(err)

	var entryNames []string
	for _, entry := range directoryEntries {
		if !entry.IsDir() {
			entryNames = append(entryNames, entry.Name())
		}
	}

	return entryNames
}

func printEntries(entries []string) {
	for _, entry := range entries {
		fmt.Println(entry)
	}
}
