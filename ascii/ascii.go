package ascii

import (
	"fmt"
	"os"
)

// AsciiArt processes words, printing their ASCII art
// character by character and adding new lines as needed.
func AsciiArt(words []string, contents2 []string, outputfilename string) {
	countSpace := 0
	finalstr := ""

	outputfile, err := os.Create(outputfilename)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer outputfile.Close()

	for _, word := range words {
		if word != "" {
			for i := 0; i < 8; i++ {
				for j, char := range word {
					if char == '\n' {
						_, err := outputfile.WriteString("\n")
						if err != nil {
							fmt.Println("Error: ", err)
						}
						continue
					}
					if !(char >= 32 && char <= 126) {
						fmt.Println("Error: Input contains non-ASCII characters")
						return
					}
					// Print the calculated index of 'char' Ascii Art in content2.
					finalstr = contents2[int(char-' ')*9+1+i]
					if j == len(word)-1{
						_, err := outputfile.WriteString(finalstr + "\n")
						if err != nil {
							fmt.Println("Error: ", err)
						}
					} else {
						_, err := outputfile.WriteString(finalstr) // writing the result in the output file
						if err != nil {
							fmt.Println("Error: ", err)
						}
					}
				}
			}
		} else {
			countSpace++
			if countSpace < len(words) {
				outputfile.WriteString("\n")
			}
		}
	}
}
