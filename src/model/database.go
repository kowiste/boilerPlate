package model

type DatabaseI interface {
	Create(data ModelI) (err error)
	FindOne(ModelI) (status int, err error)
	FindAll(map[string]string, FindAllRequest, ModelI, any) (status int, count int64, err error)
	Update(ModelI, map[string]any) (status int, err error)
	Delete(data ModelI) (status int, err error)
}
