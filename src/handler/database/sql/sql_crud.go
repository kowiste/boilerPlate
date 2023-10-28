package sql

import (
	"errors"
	"net/http"

	"serviceX/src/model"
)

const (
	ErrNotFound   string = "not Found"
	ErrValidation string = "validation error"
)

func (s db) Create(data model.ModelI) error {
	return s.conn.Create(data).Error
}

// FindOne
func (s db) FindOne(data model.ModelI) (status int, err error) {
	res := s.conn.Where("id = ?", data.GetID()).First(data)
	if res.RowsAffected == 0 {
		return http.StatusNotFound, errors.New(ErrNotFound)
	}
	return http.StatusOK, nil
}

// FindAll
func (s db) FindAll(filters map[string]string, request model.FindAllRequest, modelType model.ModelI, data any) (status int, count int64, err error) {
	query := s.conn.Model(modelType)

	if len(filters) > 0 {
		query = query.Where(filters)
	}

	query.Count(&count)

	res := query.
		Offset(request.Offset).
		Limit(request.Limit).
		Find(data)
	if res.Error != nil {
		return http.StatusInternalServerError, 0, res.Error
	}

	return http.StatusOK, count, nil
}
func (s db) Update(modelType model.ModelI, data map[string]any) (status int, err error) {
	if !s.validSchema(data, modelType) {
		return http.StatusBadRequest, errors.New(ErrValidation)
	}
	res := s.conn.Model(modelType).Where("id = ?", modelType.GetID()).Updates(data)
	if res.RowsAffected < 1 {
		return http.StatusNotFound, errors.New(ErrNotFound)
	} else if res.Error != nil {
		return http.StatusInternalServerError, res.Error
	}
	return http.StatusBadRequest, nil
}

// Delete
func (s db) Delete(data model.ModelI) (status int, err error) {
	res := s.conn.Where("id = ?", data.GetID()).Delete(data)
	if res.RowsAffected < 1 {
		return http.StatusNotFound, errors.New(ErrNotFound)
	} else if res.Error != nil {
		return http.StatusInternalServerError, res.Error
	}
	return http.StatusOK, nil
}
