package model

import (
	"fmt"
	"restapi-lesson/internal/author/storage"
)

type sortOptions struct {
	Field, Order string
}

func NewSortOptions(field, order string) storage.SortOptions {
	return &sortOptions{
		Field: field,
		Order: order,
	}
}

func (so *sortOptions) GetOrderBy() string {
	return fmt.Sprintf("%s %s", so.Field, so.Order)
}
