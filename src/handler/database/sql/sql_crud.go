package sql

import (
	"net/http"

	"serviceX/src/model"

	"github.com/gin-gonic/gin"
)

func (s db) Create(c *gin.Context, data model.ModelInterface) {
	err := s.conn.Create(data).Error
	if err != nil {
		//controller.Log().Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusCreated, data)
}

// FindOne
func (s db) FindOne(c *gin.Context, data model.ModelInterface) {
	res := s.conn.Where("id = ?", c.Param("id")).First(data)
	if res.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, data)
}

// FindAll
func (s db) FindAll(c *gin.Context, request model.FindAllRequest, modelType model.ModelInterface, data any) {
	var count int64

	query := s.conn.Model(modelType)

	filters := c.QueryMap("filter")
	if len(filters) > 0 {
		query = query.Where(filters)
	}

	query.Count(&count)

	query.
		Offset(request.Offset).
		Limit(request.Limit).
		Find(data)

	c.JSON(200, model.FindAllResponse{
		Count: count,
		Data:  data,
	})
}
func (s db) Update(c *gin.Context, modelType model.ModelInterface, data map[string]any) {
	if !s.validSchema(data, modelType) {
		c.Status(http.StatusBadRequest)
		return
	}
	res := s.conn.Model(modelType).Where("id = ?", modelType.GetID()).Updates(data)
	if res.RowsAffected < 1 {
		c.Status(http.StatusNotFound)
		return
	} else if res.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, modelType)
}
func (s db) Delete(c *gin.Context, data model.ModelInterface) {
	res := s.conn.Where("id = ?", c.Param("id")).Delete(data)
	if res.RowsAffected < 1 {
		c.Status(http.StatusNotFound)
		return
	} else if res.Error != nil {
		c.Status(http.StatusInternalServerError)
		//controller.Log().Error(err)
		return
	}
	c.Status(http.StatusOK)
}
