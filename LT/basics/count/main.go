package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/HenryPu04/count/lines"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "name of file")
	flag.Parse()

	if filename == "" {
		fmt.Println("Please provide a filename with '-f'")
		return
	}

	// Open file
	file, err := os.Open(filename)
	// Check for errors
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	// Defer closing the file to the end of the function
	defer file.Close()

	count, err := lines.Count(file)
	if err != nil {
		fmt.Println("Error counting lines:", err)
		return
	}

	fmt.Println(filename, count)
}
