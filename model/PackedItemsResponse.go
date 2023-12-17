package model

type PackedItemsInfo struct {
	ExtraItems int
	TotalItems int
	TotalPacks int
	Packs      map[string]PackInfo
}

type PackInfo struct {
	Size int `json:"size" binding:"required"`
	Used int `json:"used" binding:"required"`
}
