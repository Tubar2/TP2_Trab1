package Input

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

//Interface for a Input type
type IInput interface {
	Extract()
}

//Input implementation for a txt file
type TextInput struct {
	File           string
	ExtractedLines chan string
}

//Implementing interface by defining extract method
func (ti *TextInput) Extract() {
	file, err := os.Open(ti.File)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var line string

	//Reading every line from text file delimited with a period
	for {
		line, err = reader.ReadString('\n')
		//Sending strings to chanel
		ti.ExtractedLines <- line
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		log.Fatal(err)
	}

	//Closing channel after all lines are extracted
	close(ti.ExtractedLines)
}

type DBLPInput struct {
	site string
}

func (db *DBLPInput) extract() {
	resp, err := http.Get(db.site)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			log.Fatal(tokenizer.Err())
		}

		if tokenType == html.StartTagToken {
			token := tokenizer.Token()

			if "title" == token.Data {
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					fmt.Println(tokenizer.Token().Data)
					break
				}
			}
		}
	}

}
