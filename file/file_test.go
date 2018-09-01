package file

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/peter-mueller/sheetmusic-index"
)

func TestReadFromFile(t *testing.T) {
	sheets, err := ReadFromFile("notenindex.md")
	if err != nil {
		log.Fatal(err)
	}

	index := index.MakeIndex(sheets)

	s := index.ToMarkdown()
	ioutil.WriteFile("index.md", []byte(s), 0666)
}
