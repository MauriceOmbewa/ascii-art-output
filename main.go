package main

import (
	"fmt"
	"os"
	"strings"

	"ascii/ascii"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("Example: go run . --output=<fileName.txt> something standard")
		return
	}
	// Get an input text from the commandline.
	inputtext := os.Args[2]

	// Check if there are any args in the commandline and if the input is also empty.
	if inputtext == "" {
		return
	}

	if inputtext == "\\a" || inputtext == "\\0" || inputtext == "\\f" || inputtext == "\\v" || inputtext == "\\r" {
		fmt.Println("Error: Non printable character ", inputtext)
		return
	}
	inputtext = strings.ReplaceAll(inputtext, "\\t", "    ")
	inputtext = strings.ReplaceAll(inputtext, "\\b", "\b")
	inputtext = strings.ReplaceAll(inputtext, "\\n", "\n")

	// Logic process for handlng the backspace.
	for i := 0; i < len(inputtext); i++ {
		indexb := strings.Index(inputtext, "\b")

		if indexb > 0 {
			inputtext = inputtext[:indexb-1] + inputtext[indexb+1:]
		}
	}

	// Split our input text to a string slice and separate with a newline.
	words := strings.Split(inputtext, "\n")

	//setting the bannerfile to be used according to user input
	banner := os.Args[3]
	bannlow := strings.ToLower(banner)
	bann := bannlow + ".txt"//outputfile.WriteString("\n")


	// Read content from one of the banner files.
	contents, err := os.ReadFile(bann)
	// Apply error handling if file is incorrect.//"strings"
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fileInfo, err := os.Stat(bann)
	if err != nil {
		fmt.Println("Error reading file information", err)
		return
	}
	fileSize := fileInfo.Size()

	if fileSize == 6623 || fileSize == 4702 ||  fileSize == 7462{
		// Split the content to a string slice and separate with newline.
		contents2 := strings.Split(string(contents), "\n")

		outputfilename := os.Args[1]
		if strings.HasPrefix(outputfilename, "--output="){
			outputfilename = outputfilename[9:]
		} else {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			return
		}
		ascii.AsciiArt(words, contents2, outputfilename)
	} else {
		fmt.Println("Error with the file size", fileSize)
		return
	}
}
