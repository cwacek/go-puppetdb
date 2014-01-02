package puppetdb

import "time"

//Summarizes the success, failure,
// and skip counts for events
type EventSummary struct {
	Certname  string
	Successes int
	Failures  int
	Noops     int
	Skips     int
	Total     int
}

type Event struct {
	Certname  string
	ReportId  string `json:"report"`
	Status    string
	Timestamp *time.Time

	Type    string `json:"resource-type"`
	Title   string `json:"resource-title"`
	Message string

	File string
	Line int
}

// Retrieve a summary of events for the given certname
// (or all events if none is provided). Aggregate by
// the specified aggregator
func SummarizeEvents(aggregator, certname string) (*EventSummary, error) {
	es := new(EventSummary)
	p := NewQuery()
	p.Add("summarize-by", aggregator)

	conditions := []Condition{
		Condition{
			">",
			"timestamp",
			time.Now().AddDate(0, 0, -7).Format(time.RFC3339),
		},
	}

	if certname != "" {
		conditions = append(conditions,
			Condition{
				"=", "certname", certname,
			})
		es.Certname = certname
	}

	p.AddConditions("query", "and", conditions)

	err := api_GET(&es, "/v3/aggregate-event-counts?"+p.Encode())
	if err != nil {
		return nil, err
	}

	return es, nil
}

func EventsGetByConditions(conds ConditionSet) ([]Event, error) {
	e := make([]Event, 0)
	p := NewQuery()
	conds = append(conds,
		Condition{">", "timestamp",
			time.Now().AddDate(0, 0, -7).Format(time.RFC3339)},
	)

	p.AddConditions("query", "and", conds)

	err := api_GET(&e, "/v3/events?"+p.Encode())
	if err != nil {
		return nil, err
	}

	return e, nil
}

func EventsGetByStatus(status string) (*Event, error) {

	e := new(Event)
	p := NewQuery()
	p.AddCondition("query",
		"and",
		ConditionString(">", "timestamp",
			time.Now().AddDate(0, 0, -7).Format(time.RFC3339)),
		ConditionString("=", "status", status))

	err := api_GET(&e, "/v3/events?"+p.Encode())
	if err != nil {
		return nil, err
	}

	return e, nil
}
