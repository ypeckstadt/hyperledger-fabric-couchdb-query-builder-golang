package couchdbquerybuilder

// GreaterThanCondition The field is greater than to the argument.
type GreaterThanCondition struct {
	Value int `json:"$gt"`
}

// GreaterThanCondition The field is greater than or equal to the argument.
type GreaterThanOrEqualCondition struct {
	Value int `json:"$gte"`
}

// LessThanCondition The field is less than the argument
type LessThanCondition struct {
	Value int `json:"$lt"`
}

// LessThanOrEqualCondition The field is less than or equal to the argument.
type LessThanOrEqualCondition struct {
	Value int `json:"$lte"`
}

// EqualCondition The field is equal to the argument
type EqualCondition struct {
	Value interface{} `json:"$eq"`
}

// NotEqualCondition The field is not equal to the argument
type NotEqualCondition struct {
	Value interface{} `json:"$neq"`
}

// ExistCondition Check whether the field exists or not, regardless of its value.
type ExistCondition struct {
	Value bool `json:"$exists"`
}

// SizeCondition Special condition to match the length of an array field in a document. Non-array fields cannot match this condition.
type SizeCondition struct {
	Value uint `json:"$size"`
}

// InCondition The document field must exist in the list provided.
type InCondition struct {
	Value []interface{} `json:"$in"`
}

// TypeCondition Check the document field’s type. Valid values are "null", "boolean", "number", "string", "array", and "object".
type TypeCondition struct {
	Value string `json:"$type"`
}

// RegExCondition A regular expression pattern to match against the document field. Only matches when the field is a string value and matches the supplied regular expression. The matching algorithms are based on the Perl Compatible Regular Expression (PCRE) library. For more information about what is implemented, see the see the Erlang Regular Expressio
type RegExCondition struct {
	Value string `json:"$regex"`
}

// ModCondition [Divisor, Remainder]	Divisor and Remainder are both positive or negative integers. Non-integer values result in a 404. Matches documents where field % Divisor == Remainder is true, and only when the document field is an integer.
type ModCondition struct {
	Value [2]int `json:"$mod"`
}

