package application

import "net/url"

type TenantFilter struct {
	Name     string
	LastName string
}

func NewTenantFilterFromQueryParams(queryParams *url.Values) (*TenantFilter, error) {
	return &TenantFilter{Name: queryParams.Get("name"), LastName: queryParams.Get("lastName")}, nil
}

func (filter *TenantFilter) IsEmpty() bool {
	return filter.Name == "" && filter.LastName == ""
}
