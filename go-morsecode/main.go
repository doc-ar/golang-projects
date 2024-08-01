package main

import (
	"bufio"
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

func getSentenceMorseCode(sentence string) string {
	var morseCode string
	for _, character := range sentence {
		morseCode += getMorseCode(string(character)) + " "
	}
	return morseCode
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
	reader := bufio.NewReader(os.Stdin)
	jsonFile, err := os.ReadFile("morse.json")
	checkError(err)
	_ = json.Unmarshal(jsonFile, &morseCodes)

	fmt.Print("Welcome to Morse Code Generator. Press ctrl+c to exit.")

	fmt.Println("\nLength of Arr: ", len(morseCodes))

	var sentence string
	for sentence != "_" {
		fmt.Printf("\n\nEnter a sentence: ")
		sentence, _ := reader.ReadString('\n')

		morse := getSentenceMorseCode(sentence)

		fmt.Printf("Your morse code is: %s", morse)
	}
}
