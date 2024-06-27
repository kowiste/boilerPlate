package asset

type Asset struct {
	ID          string `json:"id"`
	ParentID    string `json:"parentID"`
	Description string `json:"description"`
}

type Assets []Asset
