package helloworld

import "fmt"

const (
	spanish = "Spanish"
	portuguese = "Portuguse"
	
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	portugueseHelloPrefix = "Oi, "
)


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language{
		case spanish: 
			prefix = spanishHelloPrefix
		case portuguese:
			prefix = portugueseHelloPrefix
		default:
			prefix = englishHelloPrefix	
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
