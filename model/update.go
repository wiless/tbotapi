package model

import "sort"

// UpdateResponse represents the response sent by the API for a GetUpdates request
type UpdateResponse struct {
	BaseResponse
	Update []Update `json:"result"`
}

// ByID is a wrapper to sort an []Update by ID
type ByID []Update

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

// Sort sorts all the updates contained in an UpdateResponse by their ID
func (resp *UpdateResponse) Sort() {
	sort.Sort(ByID(resp.Update))
}

// Update represents an incoming update
type Update struct {
	ID      int     `json:"update_id"`
	Message Message `json:"message"`
}
