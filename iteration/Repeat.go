package main

import "fmt"

// Repeat prints a specified character for a specified
// number of times.
func Repeat(character string, max int) string {
	repeatedCharacter := ""
	for i := 0; i < max; i++ {
		repeatedCharacter += character
	}
	return repeatedCharacter
}

func main() {
	fmt.Println(Repeat("e", 8))
}
