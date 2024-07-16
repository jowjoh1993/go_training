package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// Interface naming convention: method name + 'er'
type saver interface {
	// Ambiguous type that is valid for any object that has a
	// Save() method with no arguments and returns an error
	Save() error
}

// Embedded interface
type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething("hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
}

// "Any" type example
func printSomething(value interface{}) {
	typedVal, ok := value.(int)

	if ok {
		fmt.Println("Integer: ", typedVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func getNoteData() (string, string) {
	// Get input from user
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

// Get string input from user
// Return error if input is blank
func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

// Function that takes an argument of type 'saver', which we defined
// above. Valid types are any that have a Save() method taking zero
// arguments and returning an error object.
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}
