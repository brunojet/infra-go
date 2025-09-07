package repo

import (
	"context"

	"github.com/brunojet/infra-go/pkg/domain"
	"gorm.io/gorm"
)

// StorageObjectRepo é um wrapper com queries específicas para StorageObject
type StorageObjectRepo struct {
	*Repository[domain.StorageObject]
}

func NewStorageObjectRepo(db *gorm.DB) *StorageObjectRepo {
	return &StorageObjectRepo{Repository: NewRepository[domain.StorageObject](db)}
}

func (r *StorageObjectRepo) CreateWith(ctx context.Context, anexo *domain.StorageObject, createChild func(tx *gorm.DB, anexoID int64) error) error {
	// Usar transação para garantir atomicidade
	return r.WithTx(ctx, func(txRepo *Repository[domain.StorageObject]) error {
		// criar anexo
		if err := txRepo.Create(ctx, anexo); err != nil {
			return err
		}
		// delegar criação do child ao callback, usando a DB da transação
		if createChild != nil {
			if err := createChild(txRepo.DB().WithContext(ctx), anexo.ID); err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *StorageObjectRepo) GetWith(ctx context.Context, id int64, preloads ...string) (*domain.StorageObject, error) {
	return r.GetByID(ctx, id, preloads...)
}
