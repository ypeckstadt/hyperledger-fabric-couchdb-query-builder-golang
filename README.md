# CouchDB query builder for use with Hyperledger Fabric

```
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
```
