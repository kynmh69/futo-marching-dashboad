package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TimeTracking represents a time tracking record for practice sessions
type TimeTracking struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID         primitive.ObjectID `bson:"userId" json:"userId"`
	PracticeMenuID primitive.ObjectID `bson:"practiceMenuId,omitempty" json:"practiceMenuId,omitempty"`
	ClockIn        time.Time          `bson:"clockIn" json:"clockIn"`
	ClockOut       *time.Time         `bson:"clockOut,omitempty" json:"clockOut,omitempty"`
	Duration       *int64             `bson:"duration,omitempty" json:"duration,omitempty"` // Duration in seconds
	Notes          string             `bson:"notes" json:"notes"`
	CreatedAt      time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt      time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// CreateTimeTrackingInput represents data needed to create a new time tracking record
type CreateTimeTrackingInput struct {
	PracticeMenuID string `json:"practiceMenuId"`
	Notes          string `json:"notes"`
}

// UpdateTimeTrackingInput represents data needed to update a time tracking record
type UpdateTimeTrackingInput struct {
	ClockOut time.Time `json:"clockOut" validate:"required"`
	Notes    string    `json:"notes"`
}

// PrepareCreate sets fields needed for creating a new time tracking record
func (t *TimeTracking) PrepareCreate() {
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now
	t.ClockIn = now
}

// PrepareUpdate sets fields needed for updating a time tracking record
func (t *TimeTracking) PrepareUpdate() {
	t.UpdatedAt = time.Now()
	
	// Calculate duration if clock out is set
	if t.ClockOut != nil {
		duration := t.ClockOut.Unix() - t.ClockIn.Unix()
		t.Duration = &duration
	}
}