package main

import (
	"fmt"
	"os"
)

// TODO - use maps instead of slices since file names are unique, and would improve efficency
func main() {
	const sourceDirectory string = ""
	const destinationDirectory string = ""

	var sourceDirectoryEntries []string = readDirectory(sourceDirectory)
	printEntries(sourceDirectoryEntries)

	var destinationDirectoryEntries []string = readDirectory(destinationDirectory)
	printEntries(destinationDirectoryEntries)

	filesToBeRemoved := findFolderDiff(sourceDirectoryEntries, destinationDirectoryEntries)
	fmt.Println("Files in destination but not source:")
	printEntries(filesToBeRemoved)

	filesToBeAdded := findFolderDiff(destinationDirectoryEntries, sourceDirectoryEntries)
	fmt.Println("Files in source but not destination:")
	printEntries(filesToBeAdded)
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

func findFolderDiff(source []string, destination []string) []string {

	var missingEntries []string
	for _, entry := range destination {
		if !contains(source, entry) {
			missingEntries = append(missingEntries, entry)
		}
	}

	return missingEntries
}

func contains(slice []string, value string) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}

	return false
}
