package models

import (
	"time"

	"github.com/gofrs/uuid"
)

func (ms ModelSuite) Test_Task() {
	task := Task{}
	ms.Empty(task)
	task = Task{
		ID:        uuid.Must(uuid.NewV4()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	ms.NotEmpty(task)
	task = Task{
		ID:             uuid.Must(uuid.NewV4()),
		Description:    "",
		Completed:      true,
		CompletionDate: time.Now(),
		RequesterName:  "",
		ExecutorName:   "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
