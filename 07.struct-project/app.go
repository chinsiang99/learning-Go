package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/struct-project/note"
)

func main() {
	noteData, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	noteData.DisplayData()
	err = noteData.Save()
	if err != nil {
		fmt.Println("", "Saving the note failed...")
		return
	}
	fmt.Println("Saving the note succeeded...")
}

func getNoteData() (note.Note, error) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	noteData, err := note.New(title, content)
	return noteData, err
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	// var value string
	// fmt.Scanln(&value) // cannot handle spaces, can only include single word
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
