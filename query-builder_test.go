package couchdbquerybuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryBuilder(t *testing.T) {

	builder := New()

	root := builder.AddCombination(OR, Filter{Field:"id", Value:5},Filter{Field:"ids", Value:10})
	levelOne := root.AddCombination(AND,Filter{Field:"a", Value:1},Filter{Field:"b", Value:2})
	levelOne.AddCombination(OR, Filter{Field:"x", Value:0},Filter{Field:"x", Value:100})

	query, err := builder.
		SetDocType("items").
		SetLimit(10).
		SetSkip(20).
		AddField("fieldOne").
		AddField("fieldTwo").
		AddField("fieldThree", "fieldFour").
		AddField("fieldFive").
		AddFilter("id", []int{1,2,3}).
		AddFilter("name", "peckstadt").
		AddFilter("categories", []string{"a","b"}).
		AddFilter("singleID", 1).
		AddCondition("testFieldForCondition", GreaterThanCondition{Value:1}).
		AddCondition("testFieldForCondition", GreaterThanOrEqualCondition{Value:1}).
		AddCondition("testFieldForCondition", LessThanCondition{Value:1}).
		AddCondition("testFieldForCondition", LessThanOrEqualCondition{Value:1}).
		AddCondition("testFieldForCondition", EqualCondition{Value:"1"}).
		AddCondition("testFieldForCondition", NotEqualCondition{Value:1}).
		AddSort("docType", "desc").
		AddSort("createdAt", "desc").
		Build()
	expected := `{"fields":["fieldOne","fieldTwo","fieldThree","fieldFour","fieldFive"],"limit":10,"selector":{"$or":[{"id":5},{"ids":10},{"$and":[{"a":1},{"b":2},{"$or":[{"x":0},{"x":100}]}]}],"categories":["a","b"],"docType":"items","id":[1,2,3],"name":"peckstadt","singleID":1,"testFieldForCondition":{"$neq":1}},"skip":20,"sort":[{"docType":"desc"},{"createdAt":"desc"}]}`

	assert.NoError(t, err)
	assert.Equal(t, expected, query)
}
