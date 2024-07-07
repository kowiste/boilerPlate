package asset

type Asset struct {
	ID          string `json:"id"`
	ParentID    string `json:"parentID" validate:"uuid"`
	Description string `json:"description"`
}

type Assets []Asset

func (a Asset) TableName() string {
	return "assets"
}
