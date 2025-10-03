package main

import (
	"fmt"
	"io"
	"os"

	"github.com/shirou/gopsutil/v3/disk"
)

// TODO - use maps instead of slices since file names are unique, and would improve efficency
func main() {
	const sourceDirectory string = ""
	const destinationDirectory string = ""

	diskSizeStub()

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

// TODO handle errors better
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// TODO move these to another file
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

// TODO use full path for files, using string just for conceptual purposes
func deleteFiles(files []string) {
	for _, file := range files {
		error := os.Remove(file)
		check(error)
	}
}

func copyFiles(filesToCopy []string, destinationPath string) {
	for _, file := range filesToCopy {
		err := copyFile(file, destinationPath)
		if err != nil {
			fmt.Println("Error copying file:", file, ":", err)
		} else {
			fmt.Println("Copied file", file, " successfully")
		}
	}
}

func copyFile(fileToCopy string, destinationPath string) error {
	// check directory exists, do later

	sourceFile, err := os.Open(fileToCopy)
	check(err)

	defer sourceFile.Close()

	// will append file name at the end of the path
	destinationFile, err := os.Create(destinationPath)
	check(err)

	defer destinationFile.Close()

	_, err = io.Copy(sourceFile, destinationFile)
	check(err)

	return destinationFile.Sync()
}

// TODO - use these functions to calculate size remaining on disk for checks when copying
func diskSizeStub() {
	usageStat, err := disk.Usage("/")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Disk Path: %s\n", usageStat.Path)
	fmt.Printf("Total: %.2f GB\n", float64(usageStat.Total)/1e9)
	fmt.Printf("Free: %.2f GB\n", float64(usageStat.Free)/1e9)
	fmt.Printf("Used: %.2f GB\n", float64(usageStat.Used)/1e9)
	fmt.Printf("Used Percent: %.2f%%\n", usageStat.UsedPercent)
}
