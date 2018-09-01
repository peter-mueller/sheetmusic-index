package index

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type (
	// Sheet is a printed piece of music
	Sheet struct {
		Name  string
		Index uint
	}
)

type (
	// Index to find a Sheet by dirrent orderings
	Index struct {
		List []Sheet

		ByName map[string][]Sheet
		ByWord map[string][]Sheet
	}
)

// MakeIndex for a list of sheet music
func MakeIndex(sheets []Sheet) (index Index) {
	index.List = make([]Sheet, len(sheets))
	index.ByName = make(map[string][]Sheet)
	index.ByWord = make(map[string][]Sheet)
	copy(index.List, sheets)

	for _, sheet := range sheets {
		if len(sheet.Name) < 1 {
			sheet.Name = " "
		}
		firstChar := strings.ToUpper(sheet.Name[:1])
		index.ByName[firstChar] = append(index.ByName[firstChar], sheet)

		allWords := strings.Fields(sheet.Name)
		for _, word := range allWords {
			word = strings.ToLower(word)
			if isCommonWord(word) {
				continue
			}
			index.ByWord[word] = append(index.ByWord[word], sheet)
		}
	}
	return index
}

// ToMarkdown creates a human readable text
// representation of the index in markdwon format
func (i *Index) ToMarkdown() (s string) {
	var b strings.Builder

	fmt.Fprintln(&b, "Notenindex")
	fmt.Fprintf(&b, "erstellt am %s\n", time.Now().Format("02.01.2006 um 15:04 Uhr"))
	fmt.Fprintln(&b)

	fmt.Fprintln(&b, "# Liste")
	fmt.Fprintln(&b)
	for _, sheet := range i.List {
		fmt.Fprintf(&b, "(%d) %s\n", sheet.Index, sheet.Name)
	}

	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "# Nach Name")
	fmt.Fprintln(&b)
	for _, key := range i.SortedNames() {
		sheets := i.ByName[key]
		list := strings.Join(sheetStrings(sheets), ",")
		fmt.Fprintf(&b, "%s:\n", key)
		fmt.Fprintln(&b, "  "+list)
	}

	// fmt.Fprintln(&b)
	// fmt.Fprintln(&b, "# Nach Wort")
	// fmt.Fprintln(&b)
	// for _, word := range i.SortedWords() {
	// 	sheets := i.ByWord[word]
	// 	list := strings.Join(sheetStrings(sheets), ",")
	// 	fmt.Fprintln(&b, word+":")
	// 	fmt.Fprintln(&b, "  "+list)
	// }
	// fmt.Fprintln(&b)

	return b.String()
}

func (i *Index) SortedNames() (names []string) {
	for key := range i.ByName {
		names = append(names, key)
	}
	sort.Strings(names)
	return names
}

func (i *Index) SortedWords() (words []string) {
	for key := range i.ByWord {
		words = append(words, key)
	}
	sort.Strings(words)
	return words
}

// ToString representation of a sheet
func (s Sheet) ToString() string {
	return fmt.Sprintf("%s (%d)", s.Name, s.Index)
}

func sheetStrings(sheets []Sheet) (result []string) {
	for _, sheet := range sheets {
		result = append(result, sheet.ToString())
	}
	return result
}

func isCommonWord(word string) bool {
	for _, commonWord := range commonwords {
		if strings.ToLower(commonWord) == word {
			return true
		}
	}
	return false
}
