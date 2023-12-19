package main

import (
	"fmt"
	"maps"
)


func main() {

dictionnaire := make(map[string]string)

AjouterMot(dictionnaire, "python", "un langage d eprogrammation leger")
AjouterMot(dictionnaire, "canvas", "un logiciel de creation graphique")
AjouterMot(dictionnaire, "go", "un langage de programmation")


}

func add(dictionnaire map[string]string, mot, definition string) {
	dictionnaire[mot] = definition
}

func get(dictionnaire map[string]string, mot string) string {
	definition := dictionnaire[mot]
	return definition
}

func remove(dictionnaire map[string]string, mot string) {
	delete(dictionnaire, mot)
}

func liste(dictionnaire map[string]string) []string {
	var resultat []string
	for mot, definition := range dictionnaire {
		resultat = append(resultat, fmt.Sprintf("%s: %s", mot, definition))
	}
	return resultat
}