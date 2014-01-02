package puppetdb

import "time"

type NodeList []*Node

type Node struct {
	Name         string
	Active       bool
	Deactivated  *time.Time
	CatalogTime  *time.Time `json:"catalog_timestamp"`
	FactsTime    *time.Time `json:"facts_timestamp"`
	ReportedTime *time.Time `json:"report_timestamp"`

	Facts FactSet
}

type FactSet map[string]string

/*// Retrieve a list of facts for a given node*/
func (n *Node) LoadFacts() error {

	facts := make([]*Fact, 0)
	err := api_GET(&facts, "/v3/nodes/"+n.Name+"/facts")
	if err != nil {
		return nil
	}

	fact_dict := make(map[string]string)
	for _, fact := range facts {
		fact_dict[fact.Key] = fact.Value
	}

	n.Facts = fact_dict
	return nil
}

type Fact struct {
	Node  string `json:"certname"`
	Key   string `json:"name"`
	Value string
}
