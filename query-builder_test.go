package couchdbquerybuilder

import (
	"fmt"
	"testing"
)

func TestQueryBuilder(t *testing.T) {
	builder := New()

	builder.SetDocType("items")
	builder.SetLimit(10)
	builder.SetSkip(20)

	builder.AddField("fieldOne").AddField("fieldTwo")
	builder.AddField("fieldThree", "fieldFour").AddField("fieldFive")

	builder.AddFilter("id", []int{1,2,3}).AddFilter("name", "peckstadt").AddFilter("categories", []string{"a","b"}).AddFilter("singleID", 1)

	builder.AddCondition("testFieldForCondition", GreaterThanCondition{Value:1}).AddCondition("testFieldForCondition", GreaterThanOrEqualCondition{Value:1})
	builder.AddCondition("testFieldForCondition", LessThanCondition{Value:1})
	builder.AddCondition("testFieldForCondition", LessThanOrEqualCondition{Value:1})
	builder.AddCondition("testFieldForCondition", EqualCondition{Value:"1"})
	builder.AddCondition("testFieldForCondition", NotEqualCondition{Value:1})


	builder.AddSort("docType", "desc").AddSort("createdAt", "desc")


	root := builder.AddCombination(OR, Filter{Field:"id", Value:5},Filter{Field:"ids", Value:10})
	levelOne := root.AddCombination(AND,Filter{Field:"a", Value:1},Filter{Field:"b", Value:2})
	levelOne.AddCombination(OR, Filter{Field:"x", Value:0},Filter{Field:"x", Value:100})

	query, _ := builder.Build()
	fmt.Printf("the final query is* %s", query)
}
