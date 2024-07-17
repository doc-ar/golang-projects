package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type morsecode struct {
	Name  string `json:"name"`
	Morse string `json:"morse"`
}

var morseCodes []morsecode

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func getMorseCode(character string) string {
	for _, code := range morseCodes {
		if code.Name == character {
			return code.Morse
		}
	}
	return ""
}

func main() {
	jsonFileData, err := os.ReadFile("morse.json")
	checkError(err)
	_ = json.Unmarshal(jsonFileData, &morseCodes)

	var char string

	fmt.Printf("Enter a character: ")
	fmt.Scanf("%s", &char)

	fmt.Printf("Your morse code is: %s", getMorseCode(char))
}
