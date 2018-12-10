package couchdbquerybuilder

import (
	"encoding/json"
	"fmt"
	"errors"
)

// TODO sorting
// paging

// QueryBuilder used to generate couchDB queries
type QueryBuilder struct {
	ids   []string
	fields []string
	docType string
	filters map[string]interface{}
	conditions map[string]interface{}
	hasSelector bool
}

// New create a new instance of the QueryBuilder
func New() *QueryBuilder {
	return &QueryBuilder{
		filters:make(map[string]interface{}),
		conditions:make(map[string]interface{}),
	}
}

// Build constructs the query and outputs the final result
func (builder *QueryBuilder) Build() (string,error) {


	if !builder.hasSelector {
		return "", errors.New("no doctype or filters have been added for the selector")
	}

	// Initial declaration
	queryMap := map[string]interface{}{
	}

	// add fields
	if len(builder.fields) > 0 {
		queryMap["fields"] = builder.fields
	}

	// add selector
	if builder.hasSelector {
		queryMap["selector"] = map[string]interface{}{}

		selector := queryMap["selector"]
		selectorMap, ok := selector.(map[string]interface{})
		if !ok {
			return "", errors.New("something went wrong")
		}

		// add doc type
		fmt.Println(builder.docType)
		if builder.docType != "" {
			selectorMap["docType"] = builder.docType
		}

		// add filters
		if len(builder.filters) > 0 {
			for k, v := range builder.filters {
				selectorMap[k] = v
			}
		}

		// add conditions
		if len(builder.filters) > 0 {
			for k, v := range builder.conditions {
				selectorMap[k] = v
			}
		}
	}

	bytes, err := json.Marshal(&queryMap)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}


// AddField adds a field to the couchDB query
func (builder *QueryBuilder) AddField(fields ...string) {
	builder.fields = append(builder.fields, fields...)
}


// AddFilter adds a filter to filter on in the couchDB query
func (builder *QueryBuilder) AddFilter(field string, value interface{}) {
	builder.filters[field] = value
	builder.hasSelector = true
}

// SetDocType set the main doc type for the couchDB query
func (builder *QueryBuilder) SetDocType(docType string) {
	builder.docType = docType
	builder.hasSelector = true
}

// AddCondition adds a pre-defined CouchDB condition filter to the CouchDB query
func (builder *QueryBuilder) AddCondition(field string, condition interface{}) {
	builder.conditions[field] = condition
	builder.hasSelector = true
}

