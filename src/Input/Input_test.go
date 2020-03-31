package Input

import (
	"testing"
)

func TestDBLPInput(t *testing.T) {
	site := DBLPInput{site: "https://dblp.uni-trier.de/"}

	site.extract()
}
