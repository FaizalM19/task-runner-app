package server

import (
	"encoding/json"
	"net/http"
	"task-runner/internal/auth"
	"task-runner/internal/task"
	"task-runner/models"
	"task-runner/pkg/utils"
)

// Handler for login
func LoginHandler(authService *auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginReq models.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&loginReq)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		if loginReq.Username == "" || loginReq.Password == "" {
			utils.ErrorResponse(w, http.StatusBadRequest, "Username and Password are required")
			return
		}

		token := auth.GenerateToken(loginReq.Username)

		utils.JSONResponse(w, http.StatusOK, map[string]string{"Token": token})
	}
}

// Handler for creating a task
func CreateTaskHandler(taskService *task.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := taskService.CreateTask()
		if err != nil {
			utils.ErrorResponse(w, http.StatusNotFound, "Task output not found")
			return
		}
		go taskService.RunLongTask(int64(taskID))
		utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
			"message": "Task started",
			"task_id": taskID,
		})
	}
}

// Handler for getting task status
func GetTaskStatusHandler(taskService *task.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := utils.ExtractTaskID(r.URL.Path)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		task, err := taskService.GetTaskStatus(taskID)
		if err != nil {
			utils.ErrorResponse(w, http.StatusNotFound, "Task not found")
			return
		}
		utils.JSONResponse(w, http.StatusOK, map[string]string{"status": task.Status})
	}
}

// Handler for getting completed task output
func GetTaskOutputHandler(taskService *task.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := utils.ExtractTaskID(r.URL.Path)
		if err != nil {
			utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		task, err := taskService.GetTaskOutput(taskID)
		if err != nil {
			if err != nil {
				utils.ErrorResponse(w, http.StatusNotFound, "Task not found")
				return
			}
		}
		if task.Status != "Completed" {
			utils.ErrorResponse(w, http.StatusAccepted, "Task not completed yet")
			return
		}
		utils.JSONResponse(w, http.StatusOK, map[string]string{"Task": task.Status})
	}
}

// NewServer sets up the routes for the application
func NewServer(authService *auth.AuthService, taskService *task.TaskService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", LoginHandler(authService))
	mux.Handle("/tasks", auth.AuthMiddleware(http.HandlerFunc(CreateTaskHandler(taskService))))
	mux.Handle("/tasks/status/{id}", auth.AuthMiddleware(http.HandlerFunc(GetTaskStatusHandler(taskService))))
	mux.Handle("/tasks/completed/{id}", auth.AuthMiddleware(http.HandlerFunc(GetTaskOutputHandler(taskService))))
	return mux
}
