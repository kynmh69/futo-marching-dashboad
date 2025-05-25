package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PracticeMenu represents a practice schedule for a specific date
type PracticeMenu struct {
	ID          primitive.ObjectID    `bson:"_id,omitempty" json:"id,omitempty"`
	Date        time.Time             `bson:"date" json:"date"`
	Title       string                `bson:"title" json:"title"`
	Description string                `bson:"description" json:"description"`
	Items       []PracticeMenuItem    `bson:"items" json:"items"`
	CreatedBy   primitive.ObjectID    `bson:"createdBy" json:"createdBy"`
	CreatedAt   time.Time             `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time             `bson:"updatedAt" json:"updatedAt"`
}

// PracticeMenuItem represents a single item in a practice menu
type PracticeMenuItem struct {
	StartTime   time.Time `bson:"startTime" json:"startTime"`
	EndTime     time.Time `bson:"endTime" json:"endTime"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
}

// CreatePracticeMenuInput represents data needed to create a new practice menu
type CreatePracticeMenuInput struct {
	Date        time.Time               `json:"date" validate:"required"`
	Title       string                  `json:"title" validate:"required"`
	Description string                  `json:"description"`
	Items       []CreatePracticeItemInput `json:"items" validate:"required,dive"`
}

// CreatePracticeItemInput represents data needed to create a practice menu item
type CreatePracticeItemInput struct {
	StartTime   time.Time `json:"startTime" validate:"required"`
	EndTime     time.Time `json:"endTime" validate:"required,gtfield=StartTime"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
}

// UpdatePracticeMenuInput represents data needed to update an existing practice menu
type UpdatePracticeMenuInput struct {
	Date        time.Time               `json:"date"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Items       []UpdatePracticeItemInput `json:"items" validate:"omitempty,dive"`
}

// UpdatePracticeItemInput represents data needed to update a practice menu item
type UpdatePracticeItemInput struct {
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime" validate:"omitempty,gtfield=StartTime"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

// PrepareCreate sets fields needed for creating a new practice menu
func (p *PracticeMenu) PrepareCreate(userID primitive.ObjectID) {
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
	p.CreatedBy = userID
}

// PrepareUpdate sets fields needed for updating a practice menu
func (p *PracticeMenu) PrepareUpdate() {
	p.UpdatedAt = time.Now()
}