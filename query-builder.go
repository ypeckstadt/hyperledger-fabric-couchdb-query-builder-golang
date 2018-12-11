package couchdbquerybuilder

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// QueryBuilder used to generate couchDB queries
type QueryBuilder struct {
	ids   []string
	fields []string
	docType string
	filters map[string]interface{}
	conditions map[string]interface{}
	combinations []*Combination
	sort []map[string]string
	hasSelector bool
	hasLimit bool
	hasSkip bool
	limit int
	skip int
}

// Filter used to filter on a single field
type Filter struct {
	Field string
	Value interface{}
}

// Combination used for and,or,nor,all operators
type Combination struct {
	Type CombinationType
	Value []interface{}
	Filters []Filter
	Conditions []interface{}
	Combinations []*Combination
	Builder *QueryBuilder
}

// CombinationType type of combination
type CombinationType string

const (
	// Matches if all the selectors in the array match.
	AND   CombinationType = "$and"
	// Matches if any of the selectors in the array match. All selectors must use the same index.
	OR   CombinationType = "$or"
	// Matches an array value if it contains all the elements of the argument array.
	ALL   CombinationType = "$all"
	// Matches if none of the selectors in the array match.
	NOR   CombinationType = "$nor"
)

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
	queryMap := map[string]interface{}{}

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
		if len(builder.conditions) > 0 {
			for k, v := range builder.conditions {
				selectorMap[k] = v
			}
		}

		// combinations
		if len(builder.combinations) > 0 {
			for _, combination := range builder.combinations {
				addCombinationToRoot(selectorMap, combination)
			}
		}
	}

	// add sort
	if len(builder.sort) > 0 {
		queryMap["sort"] = builder.sort
	}

	// add limit
	if builder.hasLimit {
		queryMap["limit"] = builder.limit
	}

	// add skip
	if builder.hasSkip {
		queryMap["skip"] = builder.skip
	}

	bytes, err := json.Marshal(&queryMap)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}


// AddField adds a field to the couchDB query
func (builder *QueryBuilder) AddField(fields ...string) *QueryBuilder {
	builder.fields = append(builder.fields, fields...)
	return builder
}


// AddFilter adds a filter to filter on in the couchDB query
func (builder *QueryBuilder) AddFilter(field string, value interface{}) *QueryBuilder {
	builder.filters[field] = value
	builder.hasSelector = true
	return builder
}

// SetDocType set the main doc type for the couchDB query
func (builder *QueryBuilder) SetDocType(docType string) *QueryBuilder {
	builder.docType = docType
	builder.hasSelector = true
	return builder
}

// SetLimit sets the limit for paging
func (builder *QueryBuilder) SetLimit(limit int) *QueryBuilder {
	builder.limit = limit
	builder.hasLimit = true
	return builder
}

// SetSkip sets the skip value for paging
func (builder *QueryBuilder) SetSkip(skip int) *QueryBuilder {
	builder.skip = skip
	builder.hasSkip = true
	return builder
}

// AddCondition adds a pre-defined CouchDB condition filter to the CouchDB query
func (builder *QueryBuilder) AddCondition(field string, condition interface{}) *QueryBuilder {
	builder.conditions[field] = condition
	builder.hasSelector = true
	return builder
}

// AddSort adds a field to sort on in the couchDB query
func (builder *QueryBuilder) AddSort(field string, sortOrder string) *QueryBuilder {
	sort :=  map[string]string{}
	sort[field] = strings.ToLower(sortOrder)
	builder.sort = append(builder.sort, sort)
	return builder
}

// AddCombination adds a combination to the builder query
func (builder *QueryBuilder) AddCombination(combinationType CombinationType, filters ...interface{}) *Combination {
	combination := Combination{Type:combinationType,Builder:builder}

	for _, filter := range filters {
		typeName := reflect.TypeOf(filter).Name()
		if typeName == "Filter" {
			original, ok := filter.(Filter)
			if ok {
				combination.Filters = append(combination.Filters, original)
			}
		} else {
			combination.Conditions =  append(combination.Conditions, filter)
		}
	}

	builder.combinations = append(builder.combinations, &combination)
	return &combination
}

// AddCombination adds a combination to an existing one for nesting
func (c *Combination) AddCombination(combinationType CombinationType, filters ...interface{}) *Combination {
	combination := Combination{Type:combinationType}

	// loop through filters and check type, if filter then create custom, if condition add as condition
	for _, filter := range filters {
		typeName := reflect.TypeOf(filter).Name()
		if typeName == "Filter" {
			original, ok := filter.(Filter)
			if ok {
				combination.Filters = append(combination.Filters, original)
			}
		} else {
			combination.Conditions =  append(combination.Conditions, filter)
		}
	}

	c.Combinations = append(c.Combinations, &combination)

	return &combination
}


// addCombinationToRoot converts the combination to a format that is useable by CouchDB
func addCombinationToRoot(root map[string]interface{}, combination *Combination) {

	combinationType := string(combination.Type)
	combinationRoot := map[string][]interface{}{}

	//loop through filters
	for _, filter := range combination.Filters {
		filterMap := map[string]interface{}{}
		filterMap[filter.Field] = filter.Value
		combinationRoot[combinationType] = append(combinationRoot[combinationType], filterMap)
	}

	// loop through combinations
	if len(combination.Combinations) > 0 {
		fmt.Println("has child combinations")
	}
	for _, child := range combination.Combinations {
		combinationRoot[combinationType] = addCombinationToParent(combinationRoot[combinationType] ,child)
	}

	root[combinationType] = combinationRoot[combinationType]
}

func addCombinationToParent(parent []interface{}, combination *Combination) []interface{} {
	combinationType := string(combination.Type)
	combinationRoot := map[string][]interface{}{}

	//loop through filters
	for _, filter := range combination.Filters {
		filterMap := map[string]interface{}{}
		filterMap[filter.Field] = filter.Value
		combinationRoot[combinationType] = append(combinationRoot[combinationType], filterMap)
	}

	// loop through combinations
	for _, child := range combination.Combinations {
		combinationRoot[combinationType] = addCombinationToParent(combinationRoot[combinationType], child)
	}
	parent = append(parent, combinationRoot)
	return parent
}




// TODO
//$not	Selector	Matches if the given selector does not match.
//$elemMatch	Selector	Matches and returns all documents that contain an array field with at least one element that matches all the specified query criteria.
//$allMatch	Selector	Matches and returns all documents that contain an array field with all its elements matching all the specified query criteria.