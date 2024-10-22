package entity

import "github.com/google/uuid"

type Tag struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"type:varchar(255);not null;unique"`
}
