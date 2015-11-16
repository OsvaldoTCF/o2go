package models
// Gorm Model struct type

import (
	"github.com/gernest/utron"
	"time"
)

type Model struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}

func init() {
	utron.RegisterModels(&Model{})
}