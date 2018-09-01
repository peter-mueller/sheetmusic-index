package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	index "github.com/peter-mueller/sheetmusic-index"
	"github.com/peter-mueller/sheetmusic-index/file"
)

type (
	// Args for command line interface
	Args struct {
		Path   string
		Output string
	}
)

func main() {
	var args Args
	args.Path = inputPath()
	flag.StringVar(&args.Output, "output", "index.md", "output file for text formatted index")
	flag.Parse()

	sheets, err := file.ReadFromFile(args.Path)
	if err != nil {
		log.Fatal(err)
	}

	index := index.MakeIndex(sheets)

	s := index.ToMarkdown()
	ioutil.WriteFile(args.Output, []byte(s), 0666)

	log.Printf("done. %d sheets were parsed\n", len(index.List))
}

func inputPath() string {
	if len(os.Args) < 2 {
		log.Println("using [noten.md] as input file")
		return "noten.md"
	}
	return os.Args[1]
}
