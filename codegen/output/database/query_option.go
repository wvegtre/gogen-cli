package database

import (
	"strings"

	"gorm.io/gorm"
)

type QueryFilter struct {
	Field string
	Value interface{}
}

type queryFilters []QueryFilter

func (fs queryFilters) toQueryMap() map[string]interface{} {
	m := make(map[string]interface{}, 0)
	for _, v := range fs {
		m[v.Field] = v.Value
	}
	return m
}

type queryOptionArgs struct {
	Limit           int
	Page            int
	OrFilterArgs    []QueryFilter // Default query where args use And filter, if you need or filter, set OrArgs please.
	NotFilterArgs   []QueryFilter // Filter some record with value
	NotInFilterArgs []QueryFilter // Filter some record with array
	Orders          []string      // Support multiple fileds order, like "order by age desc, name"
	SelectFields    []string      // If empty, return all db fileds by default.
}

func (a queryOptionArgs) setArgsToQuery(query *gorm.DB) *gorm.DB {
	// return special feilds in result
	if len(a.SelectFields) > 0 {
		query = query.Select(a.SelectFields)
	}
	if len(a.OrFilterArgs) > 0 {
		query = query.Or(queryFilters(a.OrFilterArgs).toQueryMap())
	}
	if len(a.NotFilterArgs) > 0 {
		query = query.Not(queryFilters(a.NotFilterArgs).toQueryMap())
	}
	if len(a.NotInFilterArgs) > 0 {
		query = query.Not(queryFilters(a.NotInFilterArgs).toQueryMap())
	}
	// order result
	if len(a.Orders) > 0 {
		query = query.Order(strings.Join(a.Orders, ", "))
	}
	// limit result is required, or maybe span a full table, this is unexpect.
	query = query.Limit(a.Limit).Offset((a.Page - 1) * a.Limit)

	return query
}

type QueryOption func(op *queryOptionArgs)

func (o Operation) WithQueryLimit(limit int) QueryOption {
	return func(op *queryOptionArgs) {
		op.Limit = limit
	}
}

func (o Operation) WithQueryPage(page int) QueryOption {
	return func(op *queryOptionArgs) {
		op.Page = page
	}
}

func (o Operation) WithOrderDesc(orderBy string) QueryOption {
	return func(op *queryOptionArgs) {
		op.Orders = append(op.Orders, orderBy+" desc")
	}
}

func (o Operation) WithOrderASC(orderBy string) QueryOption {
	return func(op *queryOptionArgs) {
		op.Orders = append(op.Orders, orderBy+" asc")
	}
}

func (o Operation) WithSelectFields(fileds []string) QueryOption {
	return func(op *queryOptionArgs) {
		op.SelectFields = fileds
	}
}

func (o Operation) WithQueryOrFilter(fileds []QueryFilter) QueryOption {
	return func(op *queryOptionArgs) {
		op.OrFilterArgs = fileds
	}
}

func (o Operation) WithQueryNotFilter(fileds []QueryFilter) QueryOption {
	return func(op *queryOptionArgs) {
		op.NotFilterArgs = fileds
	}
}

func (o Operation) WithQueryNotInFilter(fileds []QueryFilter) QueryOption {
	return func(op *queryOptionArgs) {
		op.NotInFilterArgs = fileds
	}
}
