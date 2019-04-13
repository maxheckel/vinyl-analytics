package models

type UserData struct {
	InCollection bool `json:"in_collection"`
	InWantlist   bool `json:"in_wantlist"`
}
