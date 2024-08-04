package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Recursive function to traverse the JSON tree and print "type" values
func transverseAST(data interface{}, writer io.Writer) {
	switch v := data.(type) {
	case map[string]interface{}: // Map
		if typeValue, ok := v["type"]; ok {
			switch typeValue {
			case "VariableDeclaration":
				if name, ok := v["name"]; ok {
					if value, ok := v["value"]; ok {
						fmt.Fprintf(writer, "%v := \"%v\"\n", name, value)
					}
				}
			case "FunctionDeclaration":
				if name, ok := v["name"]; ok {
					fmt.Fprintln(writer, "func", name, "() {")
					for _, value := range v {
						transverseAST(value, writer)
					}
					fmt.Fprintln(writer, "}")
				}
			case "FunctionCall":
				if name, ok := v["name"]; ok {
					if name == "fmt.println" {
						name = "fmt.Println"
					}
					fmt.Fprint(writer, name, "(")
					for _, value := range v {
						transverseAST(value, writer)
						fmt.Fprint(writer, ", ")
					}
					fmt.Fprintln(writer, ")")
				}
			case "VariableCall":
				if name, ok := v["name"]; ok {
					fmt.Fprint(writer, name)
				}
			case "StringLiteral":
				if value, ok := v["value"]; ok {
					fmt.Fprintf(writer, "\"%v\"", value)
				}
			case "ImportDeclaration":
				if value, ok := v["value"]; ok {
					fmt.Fprintf(writer, "import \"%v\"\n", value)
				}
			case "Program":
				fmt.Fprintln(writer, "package main")
				for _, value := range v {
					transverseAST(value, writer)
				}

			default:
				// fmt.Fprintln(writer, typeValue)
				fmt.Println("Unknown type", typeValue)
				for _, value := range v {
					transverseAST(value, writer)
				}

			}
		}

	case []interface{}: // Array
		for _, item := range v {
			transverseAST(item, writer)
		}
	}
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return file
}

func readFileContents(file *os.File) []byte {
	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return byteValue
}

func unmarsalData(byteValue *[]byte) (data interface{}) {
	err := json.Unmarshal(*byteValue, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func main() {
	file := openFile("example-tree.json")
	defer file.Close()

	byteValue := readFileContents(file)

	data := unmarsalData(&byteValue)

	outputFile, err := os.Create("./output/output.go")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	transverseAST(data, outputFile)
}
