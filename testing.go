package main

import "fmt"

func main (){
	sentence := "Hello, World!"
	fmt.Println(sentence)
	for i, word := range sentence {
		fmt.Println(i, string(word))
	}
}
