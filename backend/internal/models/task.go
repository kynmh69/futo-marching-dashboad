package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TaskStatus represents the status of a task
type TaskStatus string

const (
	// TaskStatusTodo represents a task that is not yet started
	TaskStatusTodo TaskStatus = "todo"
	// TaskStatusInProgress represents a task that is in progress
	TaskStatusInProgress TaskStatus = "in_progress"
	// TaskStatusCompleted represents a task that is completed
	TaskStatusCompleted TaskStatus = "completed"
)

// Task represents a task in the system
type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Status      TaskStatus         `bson:"status" json:"status"`
	DueDate     *time.Time         `bson:"dueDate,omitempty" json:"dueDate,omitempty"`
	AssignedTo  primitive.ObjectID `bson:"assignedTo" json:"assignedTo"`
	CreatedBy   primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt" json:"updatedAt"`
	CompletedAt *time.Time         `bson:"completedAt,omitempty" json:"completedAt,omitempty"`
}

// CreateTaskInput represents data needed to create a new task
type CreateTaskInput struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	AssignedTo  string    `json:"assignedTo" validate:"required"`
}

// UpdateTaskInput represents data needed to update an existing task
type UpdateTaskInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status" validate:"omitempty,oneof=todo in_progress completed"`
	DueDate     *time.Time `json:"dueDate"`
	AssignedTo  string     `json:"assignedTo"`
}

// PrepareCreate sets fields needed for creating a new task
func (t *Task) PrepareCreate(userID primitive.ObjectID) {
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now
	t.Status = TaskStatusTodo
	t.CreatedBy = userID
}

// PrepareUpdate sets fields needed for updating a task
func (t *Task) PrepareUpdate() {
	t.UpdatedAt = time.Now()
	
	// If the task is being marked as completed, set the completed time
	if t.Status == TaskStatusCompleted && t.CompletedAt == nil {
		now := time.Now()
		t.CompletedAt = &now
	}
	
	// If the task was completed but is now being set back to a different status,
	// remove the completed time
	if t.Status != TaskStatusCompleted {
		t.CompletedAt = nil
	}
}