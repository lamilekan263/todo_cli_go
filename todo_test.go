package todo_test

import (
	"os"
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {

	task := "Task Name"
	l := todo.List{}

	l.Add(task)

	if l[0].Task != task {
		t.Errorf("Expects %s got %s", task, l[0].Task)
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}
	tasks := []string{
		"Task Name",
		"Task Name",
		"Task Name",
	}

	for _, item := range tasks {
		l.Add(item)
	}

	if l[0].Task != "Task Name" {
		t.Errorf("Expects %s got %s", "Task Name", l[0].Task)
	}

	if len(l) != 3 {
		t.Errorf("Expects length to be %q got %q", 3, len(l))
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expects length to be %q got %q", 2, len(l))
	}
}

func TestComplete(t *testing.T) {
	task := "Task Name"
	l := todo.List{}

	l.Add(task)

	if l[0].Done {
		t.Errorf("Expects task not to be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("Expects task  to be completed")
	}
}

func TestSave(t *testing.T) {
	task := "Task Name"
	l1 := todo.List{}
	l2 := todo.List{}

	l1.Add(task)

	file, err := os.CreateTemp("", "")

	if err != nil {
		t.Error(err)
	}

	defer os.Remove(file.Name())

	if err := l1.Save(file.Name()); err != nil {
		t.Error(err)
	}

	if err := l2.Get(file.Name()); err != nil {
		t.Error(err)
	}

	if l1[0].Task != task {
		t.Errorf("Expects %s got %s", task, l1[0].Task)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Expects %s got %s", l2[0].Task, l1[0].Task)
	}
}
