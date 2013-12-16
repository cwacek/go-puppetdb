package puppetdb

import "time"

type NodeList []*Node

type Node struct {
  Name string
  Active bool
  Deactivated *time.Time
  CatalogTime *time.Time `json:"catalog_timestamp"`
  FactsTime *time.Time `json:"facts_timestamp"`
  ReportedTime *time.Time `json:"reported_timestamp"`
}

