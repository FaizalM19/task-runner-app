package task

import (
	"log"
	"task-runner/internal/storage"
	"task-runner/models"
	"time"
)

type TaskService struct {
	storage *storage.Storage
}

func NewTaskService(storage *storage.Storage) *TaskService {
	return &TaskService{storage: storage}
}

func (taskService *TaskService) CreateTask() (int, error) {
	return taskService.storage.CreateTask()
}

func (taskService *TaskService) GetTaskStatus(taskID int) (*models.Task, error) {
	return taskService.storage.GetTaskStatus(taskID)
}

func (taskService *TaskService) GetTaskOutput(taskID int) (*models.Task, error) {
	return taskService.storage.GetTaskOutput(taskID)
}

func (taskService *TaskService) RunLongTask(taskID int64) {
	time.Sleep(1 * time.Minute)

	_, err := taskService.storage.DB.Exec("UPDATE tasks SET status = ? WHERE id = ?", "Completed", taskID)
	if err != nil {
		log.Println("Error updating task:", err)
	}
}
