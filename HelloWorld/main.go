package main

import "fmt"

const (
	english = "english"
	french  = "french"
	spanish = "spanish"

	englishHelloPrefix = "Hello"
	frenchHelloPrefix  = "Bonjour"
	spanishHelloPrefix = "Hola"
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return GreetingPrefix(lang) + ", " + name
}

func GreetingPrefix(lang string) (prefix string) {
	// var msgPrefix string

	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case english:
		prefix = englishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", "frech"))
}
