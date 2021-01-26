package main

import "fmt"

const englishhelloprefix = "Hello,"
const spanishhelloprefix = "Hola,"
const frenchhelloprefix = "Bonjour,"
const spanish = "Spanish"
const french = "French"

//this is a function
func Helloworld(name string, language string) string {
	helloprefix := englishhelloprefix
	if name == "" {
		name = "Doc"
	}
	switch language {
	case spanish:
		helloprefix = spanishhelloprefix
	case french:
		helloprefix = frenchhelloprefix
	}
	return helloprefix + name
}

func main() {
	fmt.Println(Helloworld("Saria", ""))
}
