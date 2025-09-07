package repo

import (
	"context"

	internal "github.com/brunojet/infra-go/internal/repo"
	"github.com/brunojet/infra-go/pkg/domain"
	"gorm.io/gorm"
)

type ListParams = internal.ListParams

type Repository[T any] interface {
	Create(ctx context.Context, ent *T) error
	GetByID(ctx context.Context, id int64, preloads ...string) (*T, error)
	Update(ctx context.Context, ent *T) error
	Delete(ctx context.Context, ent *T) error
	ListWithParams(ctx context.Context, p *ListParams, mod func(*gorm.DB) *gorm.DB) ([]T, int64, error)
	DB() *gorm.DB
}

func NewRepository[T any](db *gorm.DB) Repository[T] {
	return internal.NewRepository[T](db)
}

type StorageObjectRepo interface {
	CreateWith(ctx context.Context, anexo *domain.StorageObject, createChild func(tx *gorm.DB, anexoID int64) error) error
	GetWith(ctx context.Context, id int64, preloads ...string) (*domain.StorageObject, error)
	DB() *gorm.DB
}

func NewStorageObjectRepo(db *gorm.DB) StorageObjectRepo {
	return internal.NewStorageObjectRepo(db)
}
