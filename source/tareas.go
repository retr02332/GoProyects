package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) completedTask() {
	t.completed = true
}

func (t *task) updateDescription(description string) {
	t.description = description
}

func (t *task) updateName(name string) {
	t.name = name
}

func main() {
	// var t *task = &task{...}
	t := &task{
		name:        "Terminar este programa",
		description: "Programa para practicar",
	}
	fmt.Println(t)
	t.completedTask()
	t.updateDescription("Descripción actualizada")
	t.updateName("Nombre actualizado")
	fmt.Printf("%+v\n", t)
}
