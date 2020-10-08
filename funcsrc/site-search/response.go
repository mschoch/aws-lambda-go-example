package main

import (
	"fmt"
	"github.com/blugelabs/bluge/search"
)

type DocumentMatch struct {
	Document interface{}         `json:"document"`
	Score    float64             `json:"score"`
	Expl     *search.Explanation `json:"explanation"`
	ID       string              `json:"id"`
}

type AggregationValue struct {
	DisplayName string `json:"display_name"`
	FilterName  string `json:"filter_name"`
	Count       uint64 `json:"count"`
	Filtered    bool   `json:"filtered"`
}

type Aggregation struct {
	DisplayName string              `json:"display_name"`
	FilterName  string              `json:"filter_name"`
	Values      []*AggregationValue `json:"values"`
}

type SearchResponse struct {
	Query        string                  `json:"query"`
	Total        uint64                  `json:"total"`
	TopScore     float64                 `json:"top_score"`
	Hits         []*DocumentMatch        `json:"hits"`
	Duration     string                  `json:"duration"`
	Aggregations map[string]*Aggregation `json:"aggregations"`
	Message      string                  `json:"message"`
	PreviousPage int                     `json:"previousPage,omitempty"`
	NextPage     int                     `json:"nextPage,omitempty"`
}


func NewSearchResponse(query string, dmi search.DocumentMatchIterator) (*SearchResponse, error) {
	rv := &SearchResponse{
		Query: query,
	}

	next, err := dmi.Next()
	for err == nil && next != nil {
		var dm DocumentMatch
		err = next.VisitStoredFields(func(field string, value []byte) bool {
			if field == "_id" {
				dm.ID = string(value)
			}
			return true
		})
		if err != nil {
			return nil, fmt.Errorf("error visiting stored fields: %v", err)
		}
		rv.Hits = append(rv.Hits, &dm)
		next, err = dmi.Next()
	}
	if err != nil {
		return nil, fmt.Errorf("error iterating matches: %v", err)
	}

	return rv, nil
}
