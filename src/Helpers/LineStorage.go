package Helpers

import (
	"fmt"
	"strings"
)

type LineStorage struct {
	AllLines    chan string
	wordsVector []string
	shiftedVars []string
	StopWords   []string
}

// Process method will take lines read from channel while it is opened
// Later shift variations are displayed on terminal
func (storage *LineStorage) Process() {
	for {
		line, ok := <-storage.AllLines
		if ok == false {
			break
		}

		words := strings.Fields(line)
		storage.shiftedVars = makeShifts(words, storage.StopWords)

		display(storage.shiftedVars)

	}
}

func display(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
