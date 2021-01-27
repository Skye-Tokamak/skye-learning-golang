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

	return selectLanguage(language) + name
}

//公有函数大写开头，私有函数小写开头，命名返回值可以直接只写return
func selectLanguage(language string) (helloprefix string) {
	switch language {
	case spanish:
		helloprefix = spanishhelloprefix
	case french:
		helloprefix = frenchhelloprefix
	default:
		helloprefix = englishhelloprefix
	}
	return
}
func main() {
	fmt.Println(Helloworld("Saria", ""))
}
