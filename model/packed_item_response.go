package model

type PackedItemsInfo struct {
	ExtraItems int                 `json:"extra_items" binding:"required"`
	TotalItems int                 `json:"total_items" binding:"required"`
	TotalPacks int                 `json:"total_packs" binding:"required"`
	Packs      map[string]PackInfo `json:"packs" binding:"required"`
}

type PackInfo struct {
	Size int `json:"size" binding:"required"`
	Used int `json:"used" binding:"required"`
}
