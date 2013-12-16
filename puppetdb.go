package puppetdb

import "fmt"
import "net/http"
import "encoding/json"
import "errors"

var (
  uri string
  client *http.Client
)

// Configure a DBService to talk to the PuppetDB
// database
func Configure(host string, port int) {
  uri = fmt.Sprintf("http://%s:%d", host, port)
  client = &http.Client{}
}


// Return a list of all available nodes
func List() ([]*Node, error) {

  var (
  err error
  req *http.Request
  resp *http.Response
)

  req, err = http.NewRequest("GET", uri + "/v3/nodes", nil)
  if err != nil {
    return nil, err
  }

  req.Header.Add("Accept", "application/json")

  resp, err = client.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    return nil, errors.New(resp.Status)
  }

  nodes := make([]*Node, 0)
  decoder := json.NewDecoder(resp.Body)
  err = decoder.Decode(&nodes)
  if err != nil {
    return nil, err
  }

  return nodes, nil
}
