package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	CreatedAt int64  `json:"created_at"`
}
