package models

type TeamRequest struct {
	Id           string `json:"id"`
	CollectionID string `json:"collectionID"`
	TeamID       string `json:"teamID"`
	Title        string `json:"title"`
	Request      string `json:"request"`
}

type TeamRequests []TeamRequest

func (t TeamRequests) FindByCollectionId(collectionId string) []TeamRequest {
	if len(t) == 0 {
		return t
	}
	var ptrs = make([]TeamRequest, 0)

	for _, item := range t {
		if item.CollectionID == collectionId {
			ptrs = append(ptrs, item)
		}
	}

	return ptrs
}
