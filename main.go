package main

import "fmt"

func xor(input string) string {
	var result string
	for _, char := range input {
		result += fmt.Sprintf("%x", char^5)
	}
	return result
}

func main() {
	fmt.Println(xor("babe nemz is way better than you")) // add your text here
}
