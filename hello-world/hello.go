package main

import (
	"fmt"
)

const (
	spanish   = "Spanish"
	french    = "French"
	brazilian = "Brazilian"

	englishHelloPrefix   = "Hello, "
	spanishHelloPrefix   = "Hola, "
	frenchhHelloPrefix   = "Bonjour, "
	brazilianHelloPrefix = "Olá, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchhHelloPrefix
	case brazilian:
		prefix = brazilianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
