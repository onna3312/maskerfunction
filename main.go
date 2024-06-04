package main

import (
	"bufio"
	"fmt"
	"os"
)

func maskInputLink(input string) string {
	sliceFromString := []byte(input)
	var marker string = "http://"
	markerSlice := []byte(marker)
	lengthCounter := 0

	for i := 0; i < len(sliceFromString); i++ {
		if sliceFromString[i] == markerSlice[0] {
			lengthCounter = 1
			for j := 1; j < len(markerSlice) && i+j < len(sliceFromString); j++ {
				if sliceFromString[i+j] != markerSlice[j] {
					lengthCounter = 0
					break
				}
				lengthCounter++
			}
			if lengthCounter == len(markerSlice) {
				for j := i + len(markerSlice); j < len(sliceFromString) && sliceFromString[j] != ' '; j++ {
					sliceFromString[j] = '*'
				}
			}
		}
	}

	return string(sliceFromString)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your message: ")
	txt, _ := reader.ReadString('\n')
	fmt.Println(maskInputLink(txt))
	//test
}
