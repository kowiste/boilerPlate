package sql

import (
	"errors"
	"net/http"

	"serviceX/src/handler/log"
	"serviceX/src/model"

	"github.com/gin-gonic/gin"
)

const (
	ErrNotFound   string = "not Found"
	ErrValidation string = "validation error"
)

func (s db) Create(c *gin.Context, data model.ModelI) {
	err := s.conn.Create(data).Error
	if err != nil {
		log.Get().Print(log.ErrorLevel, err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, data)
}

// FindOne
func (s db) FindOne(c *gin.Context, data model.ModelI) {
	res := s.conn.Where("id = ?", c.Param("id")).First(data)
	if res.RowsAffected == 0 {
		log.Get().Print(log.ErrorLevel, ErrNotFound)
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, data)
}

// FindAll
func (s db) FindAll(c *gin.Context, request model.FindAllRequest, modelType model.ModelI, data any) {
	var count int64

	query := s.conn.Model(modelType)

	filters := c.QueryMap("filter")
	if len(filters) > 0 {
		query = query.Where(filters)
	}

	query.Count(&count)

	res := query.
		Offset(request.Offset).
		Limit(request.Limit).
		Find(data)
	if res.Error != nil {
		log.Get().Print(log.ErrorLevel, res.Error.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, model.FindAllResponse{
		Count: count,
		Data:  data,
	})
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
func (s db) Delete(c *gin.Context, data model.ModelI) {
	res := s.conn.Where("id = ?", c.Param("id")).Delete(data)
	if res.RowsAffected < 1 {
		log.Get().Print(log.ErrorLevel, ErrNotFound)
		c.Status(http.StatusNotFound)
		return
	} else if res.Error != nil {
		log.Get().Print(log.ErrorLevel, res.Error.Error())
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
