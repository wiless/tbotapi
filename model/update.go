package model

import "sort"

type UpdateResponse struct {
	BaseResponse
	Update []Update `json:"result"`
}

type ById []Update

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }

func (resp *UpdateResponse) Sort() {
	sort.Sort(ById(resp.Update))
}

type Update struct {
	Id      int     `json:"id"`
	Message Message `json:"message"`
}
