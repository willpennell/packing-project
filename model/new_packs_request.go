package model

type NewPacksRequest struct {
	PackSizes []int `json:"pack_sizes" binding"required"`
}
