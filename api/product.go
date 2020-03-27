package api

import (
	"qiqiChat/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var p service.ProductInfo
	index := c.Query("index")
	size := c.Query("size")
	character := c.Param("characteristic")
	var err error
	p.Character, err = strconv.Atoi(character)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
	}
	p.Index, err = strconv.Atoi(index)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
	}
	p.Size, err = strconv.Atoi(size)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
	}
	res := p.GetProducts()
	c.JSON(200, res)
}
func GetProductByName(c *gin.Context) {
	var p service.ProductByName
	character := c.Param("characteristic")
	err := c.ShouldBind(&p)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
	}
	//name := c.Query("name")

	p.Character, err = strconv.Atoi(character)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
	}
	//p.Name = name
	res := p.GetProductByName()
	c.JSON(200, res)
}
func GetProductByTime(c *gin.Context) {
	var service service.ProductsByTime
	character := c.Param("characteristic")
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
	}
	ch, err := strconv.Atoi(character)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
	}
	service.Character = ch
	res := service.GetProductsByTime()
	c.JSON(200, res)
}
func ModifyProductCount(c *gin.Context) {
	var service service.ProductCount
	err := c.ShouldBind(&service)
	if err == nil {
		res := service.ModifyProductCount()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func AddProduct(c *gin.Context) {
	var service service.ProductAdd
	ch, err := strconv.Atoi(c.Param("characteristic"))
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
	}
	err = c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
	}
	service.Character = ch
	res := service.AddProduct()
	c.JSON(200, res)
}
