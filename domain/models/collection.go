package models

import "encoding/json"

// export Collection
type Collection struct {
	// this field does not in export collection
	team           Team
	teamCollection TeamCollection

	Name     string       `json:"name"`
	Folders  []Collection `json:"folders"`
	Requests []Request    `json:"requests"`
	Data     string       `json:"data"`
}

// export request
type Request map[string]interface{}

func newExportRequests(teamRequest []TeamRequest) []Request {
	var ptrs = make([]Request, 0, len(teamRequest))

	for _, item := range teamRequest {
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(item.Request), &m); err == nil && m != nil {
			ptrs = append(ptrs, m)
		}
	}

	return ptrs
}

func NewDefaultCollection(team Team, teamColl TeamCollection) Collection {
	ptr := Collection{
		team:           team,
		teamCollection: teamColl,
		Name:           teamColl.Title,
		Folders:        teamColl.Folders,
		Requests:       newExportRequests(teamColl.Requests),
		Data:           teamColl.Data.ValueOrZero(),
	}
	if ptr.Folders == nil {
		ptr.Folders = make([]Collection, 0)
	}
	if ptr.Requests == nil {
		ptr.Requests = make([]Request, 0)
	}
	return ptr
}
