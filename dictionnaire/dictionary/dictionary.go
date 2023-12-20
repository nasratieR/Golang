package dictionary

import (
	"fmt"
	"sort"
)

type Dictionnaire map[string]string

func (d Dictionnaire) Add(word, definition string) {
	d[word] = definition
}

func (d Dictionnaire) Get(word string) string {
	return d[word]
}

func (d Dictionnaire) Remove(word string) {
	delete(d, word)
}

func (d Dictionnaire) List() []string {
	var result []string
	for word, definition := range d {
		result = append(result, fmt.Sprintf("%s: %s", word, definition))
	}

	sort.Strings(result)
	return result
}