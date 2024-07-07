package response

type Basic struct {
	Filter  string `form:"filter"`
	Page    int    `form:"page"`
	PerPage int    `form:"perPage"`
}
