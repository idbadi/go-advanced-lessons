package model

import (
	"restapi-lesson/internal/author/storage"
	"restapi-lesson/pkg/api/filter"
)

type filterOptions struct {
	limit  int
	fields []filter.Field
}

func NewFilterOptions(options filter.Options) storage.FilterOptions {
	return &filterOptions{limit: options.Limit(), fields: options.Fields()}

}
