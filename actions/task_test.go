package actions

import (
	"net/http"
	"tasks_kata_03/models"
	"time"

	"github.com/gofrs/uuid"
)

func (as *ActionSuite) Test_Task_Create() {
	task := models.Task{
		ID: uuid.Must(uuid.NewV4()),
	}
	res := as.JSON("/task/create").Post(task)
	as.Equal(http.StatusCreated, res.Code)
	as.Equal(1, len(taskStorage))

	res = as.JSON("/task/create").Post(task)

	as.Equal(http.StatusCreated, res.Code)
	as.Equal(2, len(taskStorage))
}

func (as *ActionSuite) Test_Task_List() {
	taskStorage = models.TaskStorage{}
	task := models.Task{
		ID: uuid.Must(uuid.NewV4()),
	}

	res := as.JSON("/task/create").Post(task)
	as.Equal(http.StatusCreated, res.Code)

	taskStorage := models.TaskStorage{}
	res = as.JSON("/task/list").Get()
	as.Equal(http.StatusOK, res.Code)
	res.Bind(&taskStorage)
	as.Equal(1, len(taskStorage))
}

func (as *ActionSuite) Test_Task_Completed_List() {
	taskStorage = models.TaskStorage{}
	task := models.Task{
		ID:             uuid.Must(uuid.NewV4()),
		Completed:      true,
		CompletionDate: time.Now(),
		ExecutorName:   "Edwin",
		RequesterName:  "Larry",
		Description:    "Task activity",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	res := as.JSON("/task/create").Post(task)
	task = models.Task{
		ID:            uuid.Must(uuid.NewV4()),
		Completed:     false,
		ExecutorName:  "Edwin",
		RequesterName: "Larry",
		Description:   "Another Task activity",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	res = as.JSON("/task/create").Post(task)

	taskStorage := models.TaskStorage{}
	res = as.JSON("/task/list/completed").Get()
	res.Bind(&taskStorage)
	as.Equal(1, len(taskStorage))

	task = models.Task{
		ID:             uuid.Must(uuid.NewV4()),
		Completed:      true,
		CompletionDate: time.Now(),
		ExecutorName:   "Edwin",
		RequesterName:  "Larry",
		Description:    "Another Task activity",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	res = as.JSON("/task/create").Post(task)

	res = as.JSON("/task/list/completed").Get()
	res.Bind(&taskStorage)
	as.Equal(2, len(taskStorage))
}
