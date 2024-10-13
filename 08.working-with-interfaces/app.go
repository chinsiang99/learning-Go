package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/struct-project/note"
	"example.com/struct-project/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

// type outputtable interface {
// 	saver
// 	displayer
// }

type outputtable interface {
	saver
	DisplayData()
}

func main() {
	noteData, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
	toDoText := getUserInput("Todo text: ")

	todo, err := todo.New(toDoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(noteData)
	if err != nil {
		fmt.Println("", "Saving the note failed...")
		return
	}
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

// func getTodoData() string {
// 	text := getUserInput("Todo text: ")
// 	return text
// }

func outputData(data outputtable) error {
	data.DisplayData()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("", "Saving the note failed...")
		return err
	}
	fmt.Println("Saving the note succeeded...")
	return nil
}

// any value = interface{} // same like in typescript
func printSomething(value interface{}) {
	// switch value.(type) {
	// case int:
	// 	// do something when the type is int
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	// do something when the type is int
	// 	fmt.Println("float: ", value)
	// case string:
	// 	// do something when the type is int
	// 	fmt.Println(value)
	// default:
	// 	fmt.Println("hehe")
	// }

	typedVal, ok := value.(int)

	if ok {
		typedVal = typedVal + 1
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		floatVal = floatVal + 1
		return
	}

	stringVal, ok := value.(string)

	if ok {
		stringVal = stringVal + ""
		return
	}
}
