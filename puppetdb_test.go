package puppetdb

import "testing"
import "fmt"
import "net/url"

func TestList(t *testing.T) {
	Configure("localhost", 8080)

	nodes, err := ListNodes()

	if err != nil {
		t.Errorf("Failed to obtain nodelist: %v", err)
	}

	println("Found nodes: ")
	for _, n := range nodes {
		fmt.Printf("%s: %s\n", n.Name, n.FactsTime)
	}
}

func TestFacts(t *testing.T) {
	var (
		err error
	)

	nodes, err := ListNodes()

	if err != nil {
		t.Errorf("Failed to obtain nodelist: %v", err)
	}

	for _, n := range nodes {

		err = n.LoadFacts()
		if err != nil {
			t.Errorf("Failed to retrieve facts for %s: %v", n.Name, err)
		}

		fmt.Printf("Found %d facts for %s\n", len(n.Facts), n.Name)
	}
}

func TestQueryConditions(t *testing.T) {

	p := NewQuery()
	p.Add("hello", "world")

	if p.Encode() != "hello=world" {
		t.Errorf("Failed to encode querystring: %s", p.Encode())
	}

	p.AddCondition("query", "=", "certname", "foo.local")

	expected := `hello=world&query=%5B%22%3D%22%2C+%22certname%22%2C+%22foo.local%22%5D`
	if p.Encode() != expected {
		t.Errorf("Failed to encode querystring: %s != %s",
			p.Encode(), expected)
	}
}

func TestEvents(t *testing.T) {

	_, err := SummarizeEvents("resource", "")

	if err != nil {
		t.Errorf("Failed to retrieve generic events: %v", err)
	}

	_, err = SummarizeEvents("resource", "foo.local")

	if err != nil {
		uri, e := url.QueryUnescape(err.Error())
		t.Errorf("Failed to retrieve events for foo.local: %s [%v]", uri, e)
	}

}
