package main

import "fmt"

type task struct {
	name        string
	description string
}

type taskList struct {
	list []*task
}

func (tasks *taskList) appendTask(t *task) {
	tasks.list = append(tasks.list, t)
}

func (tasks *taskList) popTask(index int) {
	tasks.list = append(tasks.list[:index], tasks.list[index+1:]...)
}

func (tasks *taskList) modify(t *task) {
	t.name = "update"
}

func test(lista *taskList, t3 *task, t4 task) {
	fmt.Println(len(lista.list))
	lista.appendTask(t3)
	fmt.Println(len(lista.list))
	lista.appendTask(&t4)
	fmt.Println(len(lista.list))
	lista.popTask(1)
	fmt.Println(len(lista.list))
	lista.modify(lista.list[0])
	fmt.Println(lista.list[0])
	lista.modify(lista.list[2])
	fmt.Println(lista.list[2])
}

func main() {
	t1 := &task{
		name:        "Hacer sitio web",
		description: "Sitio web de hacking",
	}

	t2 := &task{
		name:        "Hacer ejercicio",
		description: "Baloncesto",
	}

	t3 := &task{
		name:        "Dormir",
		description: "Descansar",
	}

	t4 := task{
		name:        "Something",
		description: "Something",
	}

	lista := &taskList{
		list: []*task{
			t1, t2,
		},
	}

	test(lista, t3, t4)
}
