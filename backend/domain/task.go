package domain

import "errors"

type Task struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

func (t *Task) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	return nil
}
