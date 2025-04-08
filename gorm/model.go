package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

import "time"

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
