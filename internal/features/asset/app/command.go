package app

type CreateAssetCommand struct {
	OrgID       string
	Name        string
	Description string
	Type        string
	Properties  map[string]interface{}
}

type UpdateAssetCommand struct {
	ID          string
	OrgID       string
	Name        string
	Description string
	Properties  map[string]interface{}
}
