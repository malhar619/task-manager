# Go Task Manager with Async Workers
- A production-ready Task Management API built with Golang. This system features a non-blocking background worker that automatically completes tasks after a set delay, demonstrating high-performance concurrency patterns.

- Live API URL: https://task-manager-lhmq.onrender.com

## System Architecture
- Language: Go 1.25
- Framework: Gin Gonic (Routing & Middleware)
- Database: PostgreSQL (Hosted on Render)
- Concurrency: Go Channels & Goroutines for background processing

## Authentication: JWT (JSON Web Tokens) with Bcrypt password hashing
Postman Testing Guide
To test the live API, follow these steps in order.
### 1. Register a User
```
Method: POST
URL: https://task-manager-lhmq.onrender.com/register
Body (JSON):
{
  "email": "dev@example.com",
  "password": "securepassword123",
  "role": "admin"
}
```
### 2. Login (Get JWT Token)
```
Method: POST
URL: https://task-manager-lhmq.onrender.com/login
Body (JSON):
{
  "email": "dev@example.com",
  "password": "securepassword123"
}
Action: Copy the token value from the response.
```
### 3. Create a Task (Triggers Worker)
```
Method: POST
URL: https://task-manager-lhmq.onrender.com/tasks
Header: 
Key - Authorization 
Value - <PASTE_TOKEN_HERE>
Body (JSON):
{
  "title": "Deployment Task",
  "description": "description here"
}
Expected Result: Status is 201 Created. The task is saved as pending.
```
### 4. Verify Async Completion
```
Method: GET
URL: https://task-manager-lhmq.onrender.com/tasks
Header:
Key - Authorization 
Value - <PASTE_TOKEN_HERE>

Step A: Check immediately; status will be "pending".
Step B: Wait 2 minutes, then send the request again. Status will be "completed".
```
### 5. Delete a Task
```
Method: DELETE
URL: https://task-manager-lhmq.onrender.com/tasks/:id
Header:
Key - Authorization
Value - <PASTE_TOKEN_HERE>

1. First, run your `GET /tasks` request to find an existing Task ID
2. Replace :id in the URL with that number.
3. Click Send
'''
## Core Logic: How the Worker Works
- This project uses Channels to communicate between the API and the background worker to ensure no tasks are missed.
- The Producer: When /tasks is called, the TaskController saves the task to DB and sends the Task ID into a global TaskChannel.
- The Consumer: A background Goroutine (started at app launch) stays open, waiting for IDs to arrive in the channel.
- The Delay: Once the worker receives an ID, it starts a timer (non-blocking).
- The Update: After the timer expires, it uses GORM to update the status in the PostgreSQL database. (2min)

## Project Structure:
```
├── cmd/                # Entry point
├── config/             # Database & Env configuration
├── controllers/        # Request/Response handling
├── middleware/         # JWT Authentication
├── models/             # GORM Database schemas
├── services/           # Background worker & Channel logic
├── Dockerfile          # Multi-stage build
```
