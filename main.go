package main

import "fmt"

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
	fmt.Println(changeInputLink("http://nilsoy.com asdad http://nilsoy.com 123123123 http://3"))
}
