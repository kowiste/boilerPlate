package model

type FindAllRequest struct {
	Limit  int `form:"limit" binding:"max=50,required"`
	Offset int `form:"offset"`
}

type FindAllResponse struct {
	Count int64 `json:"count" example:"1"`
	Data  any   `json:"data"`
}
