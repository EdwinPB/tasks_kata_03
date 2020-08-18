package actions

import (
	"net/http"
	"strconv"
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

// TaskCompletedList default implementation.
func TaskCompletedList(c buffalo.Context) error {
	taskCompleted := models.TaskStorage{}
	for _, task := range taskStorage {
		if task.Completed {
			taskCompleted = append(taskCompleted, task)
		}
	}
	return c.Render(http.StatusOK, r.JSON(taskCompleted))
}

// TaskNotCompletedList default implementation.
func TaskNotCompletedList(c buffalo.Context) error {
	taskNotCompleted := models.TaskStorage{}
	for _, task := range taskStorage {
		if !task.Completed {
			taskNotCompleted = append(taskNotCompleted, task)
		}
	}
	return c.Render(http.StatusOK, r.JSON(taskNotCompleted))
}

// TaskCompletedRangeList default implementation.
func TaskCompletedRangeList(c buffalo.Context) error {
	taskCompleted := models.TaskStorage{}
	from, err := strconv.Atoi(c.Param("from"))
	if err != nil {
		return err
	}
	to, err := strconv.Atoi(c.Param("to"))
	if err != nil {
		return err
	}

	for _, task := range taskStorage {
		if task.Completed && task.CompletionDate.Hour() >= from && task.CompletionDate.Hour() <= to {
			taskCompleted = append(taskCompleted, task)
		}
	}
	return c.Render(http.StatusOK, r.JSON(taskCompleted))
}
