package puppetdb

import "fmt"
import "bytes"
import "net/url"

type puppetQuery struct {
	url.Values
}

func NewQuery() puppetQuery {
	return puppetQuery{make(url.Values)}
}

func (p puppetQuery) AddCondition(key, op, field, value string) {
	p.Add(key,
		fmt.Sprintf(`["%s", "%s", "%s"]`,
			op, field, value))
}

// Add all of the conditions in the slice
// to the query. Combine them using the
// value of conjunct ('or' or 'and')
func (p puppetQuery) AddConditions(key, conjunct string, conditions ConditionSet) {

	condition_string := conditions.String()

	if len(conditions) > 1 {
		p.Add(key,
			fmt.Sprintf(`["%s", %s]`,
				"and", condition_string))
	} else {
		p.Add(key, condition_string)
	}

}

// Represent a condition as used by the
// PuppetDB API
type Condition struct {
	// The operator to use
	Op    string
	Field string
	Value string
}

func (c Condition) String() string {
	return fmt.Sprintf(`["%s", "%s", "%s"]`,
		c.Op, c.Field, c.Value)
}

type ConditionSet []Condition

func (c ConditionSet) String() string {

	buf := new(bytes.Buffer)

	for i, condition := range c {
		buf.WriteString(condition.String())
		if i < len(c)-1 {
			buf.WriteByte(',')
		}
	}

	return buf.String()
}

// Write a condition string as used by the
// PuppetDB API for queries
func ConditionString(op, field, value string) string {
	return fmt.Sprintf(`["%s", "%s", "%s"]`,
		op, field, value)

}
