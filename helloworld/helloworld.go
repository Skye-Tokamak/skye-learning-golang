package main

import "fmt"

const englishhelloprefix = "Hello,"
const spanishhelloprefix = "Hola,"
const frenchhelloprefix = "Bonjour,"
const spanish = "Spanish"
const french = "French"

//this is a function
func Helloworld(name string, language string) string {
	if name == "" {
		name = "Doc"
	}

	return SelectLanguage(language) + name
}
func SelectLanguage(language string) (helloprefix string) {
	switch language {
	case spanish:
		helloprefix = spanishhelloprefix
	case french:
		helloprefix = frenchhelloprefix
	default:
		helloprefix = englishhelloprefix
	}
	return helloprefix
}
func main() {
	fmt.Println(Helloworld("Saria", ""))
}
