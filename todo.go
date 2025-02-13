package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// type

type item struct {
	Task        string
	Done        bool
	CompletedAt time.Time
	CreatedAt   time.Time
}

type List []item

// add

func (l *List) Add(task string) {

	t := item{
		Task:        task,
		Done:        false,
		CompletedAt: time.Time{},
		CreatedAt:   time.Now(),
	}

	*l = append(*l, t)
}

// delete
func (l *List) Delete(i int) error {

	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("error: cant get %q", i)
	}

	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// complete
func (l *List) Complete(i int) error {
	ls := *l

	if i <= 0 || i > len(ls) {
		return fmt.Errorf("error: cant get %q", i)
	}

	ls[i-1].CompletedAt = time.Now()
	ls[i-1].Done = true
	return nil
}

// save
// to save we have to save it as json format and it take in the filename
func (l *List) Save(filename string) error {

	j, err := json.Marshal(l)
	if err != nil {
		return fmt.Errorf("error: converting file to json")
	}

	if err := os.WriteFile(filename, j, 0644); err != nil {
		return fmt.Errorf("error: writing to file")
	}
	return nil
}

// Get

func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return fmt.Errorf("error: reading to file")
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""
	for k, t := range *l {
		prefix := " "
		if t.Done {
			prefix = "X "
		}
		// Adjust the item number k to print numbers starting from 1 instead of 0
		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}
