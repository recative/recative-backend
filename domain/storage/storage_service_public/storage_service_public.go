package storage_service_public

import (
	"github.com/recative/recative-backend/domain/storage/storage_model"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Service interface {
	ReadStoragesByKeysAndPermissions(keys []string, permissions []string) ([]*storage_model.Storage, error)
}

type service struct {
	db    *gorm.DB
	model storage_model.Model
}

func (s *service) ReadStoragesByKeysAndPermissions(keys []string, permissions []string) ([]*storage_model.Storage, error) {
	res := make([]*storage_model.Storage, 0, len(keys))

	for _, key := range keys {
		storage, err := s.model.ReadStorageByKey(key)
		if err != nil {
			return nil, err
		}

		if lo.Every(permissions, storage.NeedPermissions) {
			res = append(res, storage)
		}
	}

	return res, nil
}
