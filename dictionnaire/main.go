package main

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

func main() {
	dictionnaire := make(Dictionnaire)

	dictionnaire.Add("python", "un lan gage de programmation leger")
	dictionnaire.Add("canvas", "un logiciel de crétion graphique")
	dictionnaire.Add("go", "un langage de programmation")

	wordToLookup := "go"
	definition := dictionnaire.Get(wordToLookup)
	fmt.Printf("Définition de %s: %s\n", wordToLookup, definition)

	wordToRemove := "canvas"
	fmt.Printf("suppression %s du dictionnnaire.\n", wordToRemove)
	dictionnaire.Remove(wordToRemove)

	fmt.Println("liste triée des mots et de leurs définitions")
	wordList := dictionnaire.List()
	for _, entry := range wordList {
		fmt.Println(entry)
	}
}
