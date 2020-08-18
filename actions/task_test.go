package actions

import (
	"net/http"
	"tasks_kata_03/models"

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

	taskStorage := models.TaskStorage{}
	res = as.JSON("/task/list").Get()
	res.Bind(&taskStorage)
	as.Equal(1, len(taskStorage))
}
