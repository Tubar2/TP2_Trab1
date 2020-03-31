package Helpers

import "strings"

func makeShifts(words, stopWords []string) (sVars []string) {
	var isStop bool
	var lineVar = make([]string, len(words))
	copy(lineVar, words)

	for _, word := range words {
		isStop = false

		//Checking if word is a stop word
		for _, sWord := range stopWords {
			if word == sWord {
				isStop = true
				break
			}
		}
		if !isStop {
			//Appending shift
			sVars = append(sVars, strings.Join(lineVar, " "))
			//Making shift
			tempWord := word
			lineVar = lineVar[1:]
			lineVar = append(lineVar, tempWord)
		}
	}
	return
}
