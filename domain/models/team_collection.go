package models

import (
	"github.com/guregu/null/zero"
)

type TeamCollection struct {
	Id       string        `json:"id"`
	ParentID zero.String   `json:"parentID"`
	TeamID   string        `json:"teamID"`
	Title    string        `json:"title"`
	Data     zero.String   `json:"data"`
	Requests []TeamRequest `json:"-" db:"-"`
	Folders  []Collection  `json:"-" db:"-"`
}

type TeamCollections []TeamCollection

// if input empty string data will be root collection
func (t TeamCollections) FindByParentId(parentId string) []TeamCollection {
	if len(t) == 0 {
		return t
	}
	var ptrs = make([]TeamCollection, 0)

	for _, item := range t {
		if item.ParentID.ValueOrZero() == parentId {
			ptrs = append(ptrs, item)
		}
	}

	return ptrs
}
