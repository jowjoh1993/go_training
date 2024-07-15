package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}

func (note Note) Display() {
	fmt.Printf("\nYour note Titled \"%v\" has the following content:\n\n%v\n\n", note.Title, note.Content)
}

func New(Title, Content string) (Note, error) {
	if Title == "" || Content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     Title,
		Content:   Content,
		CreatedAt: time.Now(),
	}, nil
}
