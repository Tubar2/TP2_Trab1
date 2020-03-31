package main

import (
	"time"
	"tp2_trab1/src/Helpers"
	"tp2_trab1/src/Input"
)

func main() {
	ch := make(chan string)

	input := Input.TextInput{
		"src/Resources/papers.txt",
		ch}

	go input.Extract()

	lineStorage := Helpers.LineStorage{
		AllLines:  ch,
		StopWords: []string{"and", "or", "with", "is", "of", "On", "on", "to", "for"},
	}

	go lineStorage.Process()

	time.Sleep(3 * time.Second)

}
