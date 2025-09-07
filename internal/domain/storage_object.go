package domain

import (
	"errors"

	"gorm.io/gorm"
)

type ObjectStatus int16

const (
	ObjectStatusPending    ObjectStatus = 0
	ObjectStatusProcessing ObjectStatus = 10
	ObjectStatusAvailable  ObjectStatus = 20
	ObjectStatusError      ObjectStatus = 30
)

type StorageObject struct {
	BaseModel
	Path     string        `gorm:"uniqueIndex;not null;type:char(40)" json:"path"`
	Name     string        `gorm:"not null;type:varchar(40)" json:"name"`
	MimeType string        `gorm:"not null;type:varchar(100)" json:"mime_type"`
	Status   *ObjectStatus `gorm:"index;not null;default:0;constraint:check(status IN (0, 10, 20, 30))" json:"status"`
}

func (StorageObject) TableName() string { return "storage_object" }

func (s *StorageObject) BeforeCreate(tx *gorm.DB) (err error) {
	if len(s.Path) < 1 {
		return errors.New("path deve ter pelo menos 1 caracter")
	}
	if len(s.Name) < 1 {
		return errors.New("name deve ter pelo menos 1 caracter")
	}
	if len(s.MimeType) < 1 {
		return errors.New("mime_type deve ter pelo menos 1 caracter")
	}
	return nil
}
