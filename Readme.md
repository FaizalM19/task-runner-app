# Task Runner APP

This application provides an API for managing and running tasks. You can create tasks, check their statuses, and retrieve the completed task.

## Requirements

- Go
- SQLite
- `.env` file with environment variables configured


## 

1. Clone the repository:
    ```bash
    git clone https://github.com/FaizalM19/task-runner-app.git
    cd task-runner
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Run the application:
    ```bash
    go run .
    ```

   The server will start on `http://localhost:8081`.


## API Endpoints

### 1. **Login**

- **URL:** `/login`

- **Method:** `POST`

- **Request Body:**
    ```json
    {
      "username": "your-username",
      "password": "your-password"
    }
    ```
- **Response:**
    - **200 OK**: "Token" : {bearer_token}.

    - **401 Unauthorized**:"error": "Username and Password are required".

### 2. **Create Task**

- **URL:** `/tasks`

- **Method:** `POST`

- **Description:** Creates a new task and starts it in the background.

- **Response:**
    ```json
    {
        "message": "Task started",
        "task_id": "{id}"
    }
    ```

### 3. **Get Task Status**

- **URL:** `/tasks/status/{id}`

- **Method:** `GET`

- **Path Parameter:**
    - `id`: The ID of the task you want to check the status for.

- **Description:** Fetches the current status of a task.

- **Response:**

    ```json
    {
        "status": "in_progress"
    }
    ```

### 4. **Get Task Output (When Completed)**

- **URL:** `/tasks/completed/{id}`

- **Method:** `GET`

- **Path Parameter:**

    - `id`: The ID of the task you want to retrieve the output for.
- **Description:** Retrieves the output of a completed task.

- **Response:**

    ```json
    {
        "Task": "completed"
    }
    ```
    - If the task is not completed yet, the response will be:
    ```json
    {
        "error": "Task not completed yet"
    }
    ```


## Explanation:

- **Login Endpoint:** Describes the expected request body and the expected response when a login request is made.

- **Task Endpoints:** The endpoints for creating tasks, getting their statuses, and getting completed task outputs are clearly described.

- **Environment Variables:** The necessary environment variables (`PORT`, and `SECRET_KEY`) are listed to ensure proper configuration.


This README will help users set up and use the API effectively.




