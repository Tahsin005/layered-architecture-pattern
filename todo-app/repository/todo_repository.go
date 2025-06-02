package repository

import (
	"database/sql"
	"errors"

	"github.com/tahsin005/layered-based-architecture/todo-app/domain"
)

type TodoRepository interface {
    Create(todo *domain.Todo) error
    GetByID(id int) (*domain.Todo, error)
    GetAll() ([]domain.Todo, error)
    Update(todo *domain.Todo) error
    Delete(id int) error
    CreateTable() error
}

type todoRepo struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *todoRepo {
	return &todoRepo{db: db}
}

func (r *todoRepo) CreateTable() error {
	query := `CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		is_done BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := r.db.Exec(query)
	return err
}

func (r *todoRepo) Create(todo *domain.Todo) error {
	query := `INSERT INTO todos (title, description, is_done) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := r.db.QueryRow(query, todo.Title, todo.Description, todo.IsDone).Scan(&todo.ID, &todo.CreatedAt)
	return err
}

func (r *todoRepo) GetByID(id int) (*domain.Todo, error) {
	query := `SELECT id, title, description, is_done, created_at FROM todos WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var todo domain.Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone, &todo.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil
}


func (r *todoRepo) GetAll() ([]domain.Todo, error) {
	query := `SELECT id, title, description, is_done, created_at FROM todos ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []domain.Todo
	for rows.Next() {
		var todo domain.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsDone, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *todoRepo) Update(todo *domain.Todo) error {
	query := `UPDATE todos SET title = $1, description = $2, is_done = $3 WHERE id = $4`
	result, err := r.db.Exec(query, todo.Title, todo.Description, todo.IsDone, todo.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *todoRepo) Delete(id int) error {
	query := `DELETE FROM todos WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
