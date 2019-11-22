package main

import "fmt"

// languages
const spanish = "Spanish"
const english = "English"
const french = "French"
const german = "German"

// Prefixes
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const germanPrefix = "Hallo, "

func greeting(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	case german:
		prefix = germanPrefix
	default:
		prefix = englishPrefix
	}

	return
}

// Hello prints "hello" to the world or a specificed name,
// and also allows you to specifc the language.
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greeting(language) + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
