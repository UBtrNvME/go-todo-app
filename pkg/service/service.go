package service

import "github.com/UBtrNvME/go-todo-app/pkg/repository"

type Authorization interface {

}

type TodoListInterface interface {

}

type TodoItemInterface interface {

}

type Service struct {
	Authorization
	TodoItemInterface
	TodoListInterface
}

func NewService(repos *repository.Repository ) *Service {
	return &Service{}
}