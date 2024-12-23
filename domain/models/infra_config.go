package models

type InfraConfig struct {
	Id                     string `json:"id" db:"id"`
	Name                   string `json:"name" db:"name"`
	Value                  string `json:"value" db:"value"`
	Active                 bool   `json:"active" db:"active"`
	IsEncrypted            bool   `json:"isEncrypted" db:"isEncrypted"`
	LastSyncedEnvFileValue string `json:"lastSyncedEnvFileValue" db:"lastSyncedEnvFileValue"`
}
