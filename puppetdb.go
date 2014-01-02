package puppetdb

import "fmt"
import "net/http"
import "encoding/json"
import "errors"

var (
	base_uri string
	client   *http.Client
)

// Configure a DBService to talk to the PuppetDB
// database
func Configure(host string, port int) {
	base_uri = fmt.Sprintf("http://%s:%d", host, port)
	client = &http.Client{}
}

// Make a request to uri and decode the response into
// obj.
func api_GET(obj interface{}, uri string) (err error) {

	var (
		req  *http.Request
		resp *http.Response
	)

	req, err = http.NewRequest("GET", base_uri+uri, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(req.URL.RequestURI() + resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&obj)
	if err != nil {
		return err
	}

	return nil
}

//Find a node by name with facts populated
func FindNodeWithFacts(name string) (*Node, error) {

	node, err := FindNode(name)
	if err != nil {
		return nil, err
	}

	node.LoadFacts()
	return node, nil
}

//Find a node by name
func FindNode(name string) (*Node, error) {
	node := new(Node)

	err := api_GET(&node, "/v3/nodes/"+name)
	if err != nil {
		return nil, err
	}

	return node, nil
}

// Return a list of all available nodes
func ListNodes() ([]*Node, error) {

	nodes := make([]*Node, 0)

	err := api_GET(&nodes, "/v3/nodes")
	if err != nil {
		return nil, err
	}

	return nodes, nil
}
