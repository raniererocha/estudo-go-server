package todo_repository

import (
	Dto "server/src/Todo/dto"
)

var todosDb []Dto.Todo

func StoreTodo(todo Dto.CreateTodoDto) Dto.Todo {
	var newTodo Dto.Todo

	newTodo.Id = len(todosDb) + 1
	newTodo.Title = todo.Title
	newTodo.Description = todo.Description
	todosDb = append(todosDb, newTodo)

	return newTodo
}

func FindTodos() []Dto.Todo {
	return todosDb
}
