package service

import (
	"errors"

	"github.com/tahsin005/layered-based-architecture/todo-app/domain"
	"github.com/tahsin005/layered-based-architecture/todo-app/repository"
)


type TodoService interface {
	CreateTodo(todo *domain.Todo) error
	GetTodoByID(id int) (*domain.Todo, error)
	GetAllTodos() ([]domain.Todo, error)
	UpdateTodo(todo *domain.Todo) error
	DeleteTodo(id int) error
	CreateTable() error
}

type todoService struct {
    repo repository.TodoRepository
}

func NewTodoService(r repository.TodoRepository) TodoService {
    return &todoService{repo: r}
}


func (s *todoService) CreateTodo(todo *domain.Todo) error {
	if todo.Title == "" {
		return errors.New("title is required")
	}
	return s.repo.Create(todo)
}

func (s *todoService) GetTodoByID(id int) (*domain.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *todoService) GetAllTodos() ([]domain.Todo, error) {
	return s.repo.GetAll()
}

func (s *todoService) UpdateTodo(todo *domain.Todo) error {
	if todo.ID == 0 {
		return errors.New("invalid todo ID")
	}
	return s.repo.Update(todo)
}

func (s *todoService) DeleteTodo(id int) error {
	if id <= 0 {
		return errors.New("invalid ID")
	}
	return s.repo.Delete(id)
}

func (s *todoService) CreateTable() error {
	return s.repo.CreateTable()
}