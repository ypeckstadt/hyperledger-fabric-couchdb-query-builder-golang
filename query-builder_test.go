package couchdbquerybuilder

import (
	"fmt"
	"testing"
)

func TestQueryBuilder(t *testing.T) {
	builder := New()

	builder.SetDocType("kaigo-care-accident-info")

	builder.AddField("fieldOne")
	builder.AddField("fieldTwo")

	builder.AddFilter("id", []int{1,2,3})
	builder.AddFilter("name", "peckstadt")
	builder.AddFilter("categories", []string{"a","b"})
	builder.AddFilter("singleID", 1)

	builder.AddCondition("testFieldForCondition", GreaterThanCondition{Value:1})
	builder.AddCondition("testFieldForCondition", GreaterThanOrEqualCondition{Value:1})
	builder.AddCondition("testFieldForCondition", LessThanCondition{Value:1})
	builder.AddCondition("testFieldForCondition", LessThanOrEqualCondition{Value:1})
	builder.AddCondition("testFieldForCondition", EqualCondition{Value:"1"})
	builder.AddCondition("testFieldForCondition", NotEqualCondition{Value:1})

	query, _ := builder.Build()
	fmt.Printf("the final query is* %s", query)
}
