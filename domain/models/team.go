package models

type Team struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	LastUpdated string `json:"last_updated" db:"-"`
}
