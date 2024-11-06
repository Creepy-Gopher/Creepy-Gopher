package service

import (
	"context"
	"creepy/internal/filter"
)

type FilterService struct {
	filterOps *filter.Ops
	
}

func NewFilterService(filterOps *filter.Ops) *FilterService {
	return &FilterService{
		filterOps: filterOps,
	}
}

func (s *FilterService) CreateFilter(ctx context.Context, f *filter.FilterSet) error {

	if err := s.filterOps.CreateFiletr(ctx, f); err != nil {
		return err
	}

	return nil
}
