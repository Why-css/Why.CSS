package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"unicode"
)

// This is for testing purposes
func main() {

	// read .scss file
	fileContent, readErr := os.ReadFile("../main.scss")
	if readErr != nil {
		log.Fatal(readErr)
	}

	// convert .scss file content to string
	var fileContentAsString = string(fileContent)

	// convert string to list of tokens
	var listOfTokens = tokenizer(fileContentAsString)

	jsonString, jsonErr := json.Marshal(listOfTokens)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	os.WriteFile("tokenisedScss.json", jsonString, os.ModePerm)
}

func assignByteValue(byteArray []rune, value int) rune {

	if len(byteArray) < value {
		var difference = value - len(byteArray)
		return byteArray[value-difference-1]
	} else if len(byteArray) == value {
		return byteArray[value-1]
	} else {
		return byteArray[value]
	}
}

func tokenizer(input string) []map[string]string {

	var current = 0
	tokens := []map[string]string{}

	letters := regexp.MustCompile("[a-z]+")
	numbers := regexp.MustCompile("[0-9]+")

	byteArray := []rune(input)

	for current < len(byteArray) {

		char := assignByteValue(byteArray, current)

		if char == '(' {
			mapOfToken := map[string]string{"paren": "("}
			tokens = append(tokens, mapOfToken)

			current++

			continue
		}

		if char == ')' {
			mapOfToken := map[string]string{"paren": ")"}
			tokens = append(tokens, mapOfToken)

			current++

			continue
		}

		if char == '[' {
			mapOfToken := map[string]string{"brack": "["}
			tokens = append(tokens, mapOfToken)

			current++

			continue
		}

		if char == ']' {
			mapOfToken := map[string]string{"brack": "]"}
			tokens = append(tokens, mapOfToken)

			current++

			continue
		}

		if char == '{' {
			mapOfToken := map[string]string{"brace": "{"}
			tokens = append(tokens, mapOfToken)

			current++

			continue
		}

		if char == '}' {
			mapOfToken := map[string]string{"brace": "}"}
			tokens = append(tokens, mapOfToken)

			current++

			continue
		}

		if char == '@' {
			var value string

			current++

			char = assignByteValue(byteArray, current)

			for letters.MatchString(string(char)) { ///bruuuh
				value += string(char)
				current++
				char = assignByteValue(byteArray, current)
			}

			mapOfToken := map[string]string{"at": value}
			tokens = append(tokens, mapOfToken)

			continue
		}

		if char == '$' {
			var value string
			current++
			char = assignByteValue(byteArray, current)

			for letters.MatchString(string(char)) {
				value += string(char)
				current++
				char = assignByteValue(byteArray, current)
			}

			mapOfToken := map[string]string{"variable": value}
			tokens = append(tokens, mapOfToken)

			continue
		}

		if char == '=' {

			mapOfToken := map[string]string{"equal": "="}
			tokens = append(tokens, mapOfToken)

			current++

			continue
		}

		if char == '.' {

			mapOfToken := map[string]string{"dot": "."}
			tokens = append(tokens, mapOfToken)

			current++

			continue
		}

		// Check for whitespace
		if unicode.IsSpace(char) {

			current++

			continue
		}

		// Check for semicolon
		if char == ';' {

			mapOfToken := map[string]string{"end": ";"}
			tokens = append(tokens, mapOfToken)

			current++

			continue
		}

		// Check for numbers
		if numbers.MatchString(string(char)) {
			var value string

			for numbers.MatchString(string(char)) {
				value += string(char)
				current++
				char = assignByteValue(byteArray, current)
			}

			mapOfToken := map[string]string{"number": value}
			tokens = append(tokens, mapOfToken)

			continue
		}

		// Check for string
		if char == '"' {
			var value string
			current++
			char = assignByteValue(byteArray, current)

			for char != '"' {
				value += string(char)
				current++
				char = assignByteValue(byteArray, current)
			}

			mapOfToken := map[string]string{"string": value}
			tokens = append(tokens, mapOfToken)

			current++
			continue
		}

		// Check for names
		if letters.MatchString(string(char)) {
			var value string

			for char != '(' {
				value += string(char)
				current++
				char = assignByteValue(byteArray, current)
			}

			mapOfToken := map[string]string{"name": value}
			tokens = append(tokens, mapOfToken)

			continue
		}

		fmt.Println(errors.New("Don't know this character " + string(char)))

		current++
	}
	return tokens
}
