package services

import (
	"log"
	"os"
	"strconv"
	"task-manager/repository"
	"time"
)

// Channel to receive new Task IDs (The "Worker Queue")
var TaskChannel = make(chan string, 100)

func StartWorker() {
	go func() {
		for taskID := range TaskChannel {
			// Process each task in a separate goroutine to prevent blocking
			go processTask(taskID)
		}
	}()
	log.Println("Background worker started...")
}

func processTask(taskID string) {
	delayStr := os.Getenv("AUTO_COMPLETE_DELAY")
	delay, _ := strconv.Atoi(delayStr)
	if delay == 0 {
		delay = 1 // Default safety
	}

	// 1. Wait for X minutes [cite: 41]
	time.Sleep(time.Duration(delay) * time.Minute)

	// 2. Fetch fresh task data
	task, err := repository.GetTaskByID(taskID)
	if err != nil {
		return // Task might have been deleted
	}

	// 3. Check if manually completed or deleted [cite: 42]
	if task.Status == "pending" || task.Status == "in_progress" {
		repository.UpdateTaskStatus(taskID, "completed")
		log.Printf("Worker: Auto-completed Task %s", taskID)
	}
}
