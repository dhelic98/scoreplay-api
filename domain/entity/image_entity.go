package entity

import "github.com/google/uuid"

type Image struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name string    `gorm:"type:varchar(255);not null"`
	URL  string    `gorm:"type:varchar(255);not null"`
	Tags []Tag     `gorm:"many2many:image_tags;"`
}
