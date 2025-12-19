package models

import (
	"time"

	"github.com/google/uuid"
)

// task
type Task struct {
	Id          uuid.UUID
	Description string
	Status      TASKSTATUS
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// custom type
type TASKSTATUS int

// status code
const (
	TODO TASKSTATUS = iota
	IN_PROGRESS
	DONE
	UNKNOWN
)
