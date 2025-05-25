package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Event represents a calendar event or scheduled practice
type Event struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string               `bson:"title" json:"title"`
	Description string               `bson:"description" json:"description"`
	StartTime   time.Time            `bson:"startTime" json:"startTime"`
	EndTime     time.Time            `bson:"endTime" json:"endTime"`
	AllDay      bool                 `bson:"allDay" json:"allDay"`
	Attendees   []primitive.ObjectID `bson:"attendees" json:"attendees"`
	CreatedBy   primitive.ObjectID   `bson:"createdBy" json:"createdBy"`
	CreatedAt   time.Time            `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time            `bson:"updatedAt" json:"updatedAt"`
}

// Attendance represents attendance record for an event
type Attendance struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EventID   primitive.ObjectID `bson:"eventId" json:"eventId"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Status    AttendanceStatus   `bson:"status" json:"status"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

// AttendanceStatus represents the attendance status for an event
type AttendanceStatus string

const (
	// AttendancePresent means the user was present
	AttendancePresent AttendanceStatus = "present"
	// AttendanceAbsent means the user was absent
	AttendanceAbsent AttendanceStatus = "absent"
	// AttendanceExcused means the user was excused from attendance
	AttendanceExcused AttendanceStatus = "excused"
	// AttendanceLate means the user was late
	AttendanceLate AttendanceStatus = "late"
)

// CreateEventInput represents data needed to create a new event
type CreateEventInput struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime" validate:"required"`
	EndTime     time.Time `json:"endTime" validate:"required,gtfield=StartTime"`
	AllDay      bool      `json:"allDay"`
	Attendees   []string  `json:"attendees"`
}

// UpdateEventInput represents data needed to update an existing event
type UpdateEventInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime" validate:"omitempty,gtfield=StartTime"`
	AllDay      bool      `json:"allDay"`
	Attendees   []string  `json:"attendees"`
}

// PrepareCreate sets fields needed for creating a new event
func (e *Event) PrepareCreate(userID primitive.ObjectID) {
	now := time.Now()
	e.CreatedAt = now
	e.UpdatedAt = now
	e.CreatedBy = userID
}

// PrepareUpdate sets fields needed for updating an event
func (e *Event) PrepareUpdate() {
	e.UpdatedAt = time.Now()
}