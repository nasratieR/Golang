package main

import (
	"fmt"
	"dictionnaire/dictionary"
)

func main() {
	dict := dictionary.Dictionnaire{}

	dict.Add("python", "un langage de programmation léger")
	dict.Add("canvas", "un logiciel de création graphique")
	dict.Add("go", "un langage de programmation")

	motC := "go"
	definition := dict.Get(motC)
	fmt.Printf("Définition de %s: %s\n", motC, definition)

	motS := "canvas"
	fmt.Printf("Suppression de %s du dictionnaire.\n", motS)
	dict.Remove(motS)

	fmt.Println("Liste des mots et de leurs définitions :")
	motL := dict.List()
	for _, entry := range motL {
		fmt.Println(entry)
	}
}
