package model

import "github.com/gin-gonic/gin"

type Service interface {
	Authorization(c *gin.Context)
	Create(c *gin.Context)
	List(c *gin.Context)
	Find(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
