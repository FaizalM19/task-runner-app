package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"task-runner/models"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) (*sql.DB, error) {
	connStr := fmt.Sprintf("file:%s?cache=shared", dbPath)
	DB, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return nil, err
	}

	if err := DB.Ping(); err != nil {
		return nil, err
	}
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		status TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return DB, nil
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{DB: db}
}

type Storage struct {
	DB *sql.DB
}

func (s *Storage) CreateTask() (int, error) {
	statement := `INSERT INTO tasks (status, created_at, updated_at) 
				  VALUES (?, ?, ?)`
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	res, err := s.DB.Exec(statement, "in_progress", timestamp, timestamp)
	if err != nil {
		fmt.Println("enterrrr", err)
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}

func (s *Storage) GetTaskStatus(taskID int) (*models.Task, error) {
	if taskID <= 0 {
		return nil, errors.New("task ID is invalid")
	}

	var task models.Task
	err := s.DB.QueryRow("SELECT id, status FROM tasks WHERE id = ?", taskID).Scan(&task.ID, &task.Status)
	if err != nil {
		return nil, err
	}

	return &task, nil

}

func (s *Storage) GetTaskOutput(taskID int) (*models.Task, error) {
	if taskID <= 0 {
		return nil, errors.New("task ID is invalid")
	}

	var task models.Task
	err := s.DB.QueryRow("SELECT id, status FROM tasks WHERE id = ?", taskID).Scan(&task.ID, &task.Status)
	if err != nil {
		return nil, err
	}

	return &task, err
}
