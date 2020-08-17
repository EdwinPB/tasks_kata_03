package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Task struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Description    string    `json:"description" db:"description"`
	Completed      bool      `json:"completed" db:"completed"`
	CompletionDate time.Time `json:"completion_date" db:"completion_date"`
	RequesterName  string    `json:"requester_name" db:"requester_name"`
	ExecutorName   string    `json:"executor_name" db:"executor_name"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Tasks is not required by pop and may be deleted
type Tasks []Task
