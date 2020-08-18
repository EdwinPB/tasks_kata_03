package actions

import (
	"net/http"
	"tasks_kata_03/models"

	"github.com/gofrs/uuid"
)

func (as *ActionSuite) Test_Task_List() {
	as.Empty(taskCalled)
	as.JSON("/task/list").Get()
	as.NotEmpty(taskCalled)
}

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
