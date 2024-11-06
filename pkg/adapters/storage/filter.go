package storage

import (
	"context"
	"creepy/internal/filter"
	"creepy/pkg/adapters/storage/mappers"

	"gorm.io/gorm"
)

type FilterSetRepo struct {
	db *gorm.DB
}


func NewFilterSetRepo(db *gorm.DB) filter.Repo{
	return &FilterSetRepo{db}
}

func (fs *FilterSetRepo)Insert(ctx context.Context, f *filter.FilterSet)error{

	filterEntity := mappers.FilterDomainToFilterEntity(f)
	if err:= fs.db.WithContext(ctx).Save(&filterEntity).Error; err!=nil{
		return err
	}
	f.ID = filterEntity.ID
	return nil
}
