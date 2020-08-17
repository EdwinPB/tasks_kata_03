package models

// TaskStorage is not required by pop and may be deleted
type TaskStorage Tasks

// Add ..
func (ts *TaskStorage) Add(task Task) {
	*ts = append(*ts, task)
}
