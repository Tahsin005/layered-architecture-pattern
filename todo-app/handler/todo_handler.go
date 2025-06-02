package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tahsin005/layered-based-architecture/todo-app/domain"
	"github.com/tahsin005/layered-based-architecture/todo-app/service"
)

type TodoHandler struct {
	svc service.TodoService
}

func NewTodoHandler(r *mux.Router, svc service.TodoService) {
	h := &TodoHandler{svc}

	r.HandleFunc("/create-table", h.CreateTable).Methods(http.MethodGet)
	r.HandleFunc("/todos", h.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", h.GetTodoByID).Methods(http.MethodGet)
	r.HandleFunc("/todos", h.GetAllTodos).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", h.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", h.DeleteTodo).Methods(http.MethodDelete)
}

func (h *TodoHandler) CreateTable(w http.ResponseWriter, r *http.Request) {
	if err := h.svc.CreateTable(); err != nil {
		http.Error(w, "Failed to create table", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Table created successfully"))
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if err := h.svc.CreateTodo(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	todo, err := h.svc.GetTodoByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve todo", http.StatusInternalServerError)
		return
	}
	if todo == nil {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.svc.GetAllTodos()
	if err != nil {
		http.Error(w, "Failed to fetch todos", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var todo domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	todo.ID = id
	if err := h.svc.UpdateTodo(&todo); err != nil {
		http.Error(w, "Failed to update todo", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.svc.DeleteTodo(id); err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Todo deleted successfully"))
}
