package main

import (
	"fmt"

	"github.com/blugelabs/bluge"
	querystr "github.com/blugelabs/query_string"
)

const resultsPerPage = 10

type Filter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type SearchRequest struct {
	Query   string    `json:"query"`
	Filters []*Filter `json:"filters"`
	Page    int       `json:"page"`
}

func (r *SearchRequest) buildFilterClauses() (rv []bluge.Query) {
	return rv
}

func (r *SearchRequest) SizeOffset() (size, offset int) {
	return resultsPerPage, (r.Page - 1) * resultsPerPage
}

func (r *SearchRequest) BlugeRequest() (bluge.SearchRequest, error) {
	userQuery, err := querystr.ParseQueryString(r.Query, querystr.DefaultOptions())
	if err != nil {
		return nil, fmt.Errorf("errror parsing query string '%s': %v", r.Query, err)
	}

	if r.Page < 1 {
		r.Page = 1
	}

	size, offset := r.SizeOffset()

	filters := r.buildFilterClauses()

	q := bluge.NewBooleanQuery().
		AddMust(userQuery).
		AddMust(filters...)

	blugeRequest := bluge.NewTopNSearch(size, q).
		WithStandardAggregations().
		SetFrom(offset).
		ExplainScores()

	return blugeRequest, nil
}