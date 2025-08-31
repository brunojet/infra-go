package domain

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime;index" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type BaseEntity struct {
	BaseModel
	Nome      string `gorm:"not null;index" json:"nome"` // Obrigatório, index genérico para compatibilidade
	Descricao string `json:"descricao"`
	Ativo     bool   `gorm:"index;default:false" json:"ativo"`
}

func (e *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if len(e.Nome) < 1 {
		return errors.New("nome deve ter pelo menos 1 caracter")
	}
	return nil
}

type Anexo struct {
	BaseModel
	Nome          string `gorm:"size:128;not null" json:"nome"`
	TipoMime      string `gorm:"size:64;not null" json:"tipo_mime"`
	SHA256        string `gorm:"size:64;not null" json:"sha256"`
	Tamanho       int64  `gorm:"not null" json:"tamanho"`
	Armazenamento string `gorm:"size:256;not null" json:"armazenamento"`
	Caminho       string `gorm:"size:256;not null" json:"caminho"`
	Presente      bool   `gorm:"default:false" json:"presente"`
}

func (Anexo) TableName() string { return "anexo" }
