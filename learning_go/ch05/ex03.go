package main

import "fmt"

func prefixer(s string) func(string) string {
	return func(arg string) string {
		return s + " " + arg
	}
}

func main() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Ross"))
	fmt.Println(helloPrefix("Joey"))
}
