package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Text string `json:"text"`
}

func (t *Todo) IsDone() bool {
	return t.Text != ""
}
