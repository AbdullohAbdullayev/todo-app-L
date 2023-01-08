package service

import "github.com/AbdullohAbdullayev/todo-app-L.git/pkg/repository"

type Authorization interface{}
type TodoList interface{}
type TodoItem interface{}
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}