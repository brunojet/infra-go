package domain

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64     `gorm:"primaryKey:pk_base,priority:10;autoIncrement;column:id" json:"id"`
	CreatedAt time.Time `gorm:"index;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;autoUpdateTime" json:"updated_at"`
}

type BaseEntity struct {
	BaseModel
	Name        string `gorm:"index;not null" json:"name"`
	Description string `json:"description"`
	Active      bool   `gorm:"index;default:false" json:"active"`
}

func (e *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if len(e.Name) < 1 {
		return errors.New("name deve ter pelo menos 1 caracter")
	}
	return nil
}
