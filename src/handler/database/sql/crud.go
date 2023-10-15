package sql

import (
	"net/http"
	"strings"

	"serviceX/src/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

	var extraFieldsArr []string
	extraFields, ok := c.GetQuery("extraFields")
	if ok {
		extraFieldsArr = strings.Split(extraFields, ",")
		caser := cases.Title(language.English)
		for _, extraField := range extraFieldsArr {
			query = query.Preload(caser.String(extraField))
		}
	}

	filters := c.QueryMap("filter")
	if len(filters) > 0 {
		query = query.Where(filters)
	}

	/* 	if callback != nil {
		query = (*callback)(query)
	} */

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

	// Check if record exists
	res := s.conn.Where("id = ?", c.Param("id")).First(&modelType)

	if res.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	err := s.conn.Model(modelType).Where("id = ?", modelType.GetID()).Updates(data).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, modelType)
}
func (s db) Delete(c *gin.Context, data model.ModelInterface) {
	//TODO: solo una llamada y revisar error?
	res := s.conn.Where("id = ?", c.Param("id")).First(data)
	if res.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	err := s.conn.Where("id = ?", data.GetID()).Delete(data).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
		//controller.Log().Error(err)
		return
	}

	c.Status(http.StatusOK)
}
