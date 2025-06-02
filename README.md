# 📝 Go TODO App

A simple TODO application built with **Golang**, using **Gorilla Mux** for routing and **PostgreSQL** (via Docker) for data storage. It follows a clean, layered-based architecture:

- **Config Layer**
- **Database Layer**
- **Data Access Layer (Repository)**
- **Domain Layer (Entities)**
- **Application Layer (Service)**
- **Presentation Layer (Handler)**

---

## Project Structure

```
todo-app/
├── config/
│   └── config.go
├── database/
│   └── database.go
├── domain/
│   └── todo.go
├── repository/
│   └── todo_repository.go
├── service/
│   └── todo_service.go
├── handler/
│   └── todo_handler.go
├── main.go
├── go.mod
├── go.sum
└── docker-compose.yml
```


---

## 🐳 Getting Started

### 1. Clone and start PostgreSQL
```bash
git clone git@github.com:Tahsin005/layered-architecture-pattern.git
cd layered-architecture-pattern
docker compose up -d
```

---

### 2. Run the Go app
```bash
go run main.go
```

## 🧪 Tech Stack
```
Go (Golang) – core backend

Gorilla Mux – routing

PostgreSQL – database

Docker Compose – DB setup
```

## 📌 Notes
Make sure to create the table by hitting /create-table before testing CRUD endpoints.


