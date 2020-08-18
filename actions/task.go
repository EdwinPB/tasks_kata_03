package actions

import (
	"net/http"
	"tasks_kata_03/models"

	"github.com/gobuffalo/buffalo"
)

var taskStorage models.TaskStorage

// TaskCreate default implementation.
func TaskCreate(c buffalo.Context) error {
	t := models.Task{}
	if err := c.Bind(&t); err != nil {
		return err
	}
	taskStorage.Add(t)
	return c.Render(http.StatusCreated, r.JSON(t))
}

// TaskList default implementation.
func TaskList(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(taskStorage))
}
