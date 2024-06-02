package main

import (
	"bufio"
	"fmt"
	"os"
)

func changeInputLink(input string) string {
	sliceFromString := []byte(input)

	for i := 0; i < len(input); i++ {
		if input[i] == 'h' && input[i+1] == 't' && input[i+2] == 't' && input[i+3] == 'p' && input[i+4] == ':' && input[i+5] == '/' && input[i+6] == '/' {
			for x := i + 7; x < len(input); x++ {
				if sliceFromString[x] == ' ' {
					break
				}
				sliceFromString[x] = '*'

			}
		}
	}
	return string(sliceFromString)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter URL: ")
	txt, _ := reader.ReadString('\n')
	fmt.Println(changeInputLink(txt))

}
