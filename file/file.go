package file

import (
	"bufio"
	"os"

	index "github.com/peter-mueller/sheetmusic-index"
)

// ReadFromFile can parse a text file with lines as separate sheets
func ReadFromFile(path string) (sheets []index.Sheet, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var i uint
	for scanner.Scan() {
		sheets = append(sheets, index.Sheet{
			Index: i,
			Name:  scanner.Text(),
		})
		i++
	}

	return sheets, scanner.Err()
}
