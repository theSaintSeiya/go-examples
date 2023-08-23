package main

import (
	"fmt"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/files/text.txt"
	c, err := readFile(filePath)

	if err != nil {
		return
	}

	newContent := fmt.Sprintf("Original: %v\nNew: %v %v", c, c, c)

	writeErr := writeFile(filePath, newContent)
	if writeErr == nil {
		readFile(filePath)
	}
}

func readFile(path string) (string, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println(string(content))
	}

	return string(content), err
}

func writeFile(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0000)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	return err
}
