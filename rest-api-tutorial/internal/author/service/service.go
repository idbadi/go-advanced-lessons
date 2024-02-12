package service

import (
	"context"
	"fmt"
	"restapi-lesson/internal/author/model"
	"restapi-lesson/internal/author/storage"
	model2 "restapi-lesson/internal/author/storage/model"
	"restapi-lesson/pkg/api/filter"
	"restapi-lesson/pkg/api/sort"
	"restapi-lesson/pkg/logging"
)

type Service struct {
	repository storage.Repository
	logger     *logging.Logger
}

func NewService(repository storage.Repository, logger *logging.Logger) *Service {
	return &Service{repository: repository, logger: logger}
}

func (s *Service) GetAll(ctx context.Context, filterOptions filter.Options, sortOptions sort.Options) ([]model.Author, error) {
	options := model2.NewSortOptions(sortOptions.Field, sortOptions.Order)
	all, err := s.repository.FindAll(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("failed to get all authors due to error %v", err)
	}
	return all, nil
}
