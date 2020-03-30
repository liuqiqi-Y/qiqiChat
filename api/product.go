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
		return
	}
	p.Index, err = strconv.Atoi(index)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
		return
	}
	p.Size, err = strconv.Atoi(size)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
		return
	}
	res := p.GetProducts()
	//s, _ := json.Marshal(&res)
	//fmt.Println("====>" + string(s))
	c.JSON(200, res)
}
func GetProductByName(c *gin.Context) {
	var p service.ProductByName
	character := c.Param("characteristic")
	err := c.ShouldBind(&p)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	//name := c.Query("name")

	p.Character, err = strconv.Atoi(character)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
		return
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
		return
	}
	ch, err := strconv.Atoi(character)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
		return
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
		return
	}
	err = c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	service.Character = ch
	res := service.AddProduct()
	c.JSON(200, res)
}
func DelProduct(c *gin.Context) {
	var service service.ProductByName
	err := c.ShouldBind(&service) //不能绑定在请求体里的url数据
	//service.Name = c.Query("name")
	//service.Name = c.PostForm("name")

	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}

	character := c.Param("characteristic")

	ch, err := strconv.Atoi(character)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
		return
	}
	service.Character = ch
	res := service.DelProduct()
	c.JSON(200, res)
}
func ModifyProductName(c *gin.Context) {
	character := c.Param("characteristic")
	var service service.ProductName
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	ch, err := strconv.Atoi(character)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
		return
	}
	service.Character = ch
	res := service.ModifyProductName()
	c.JSON(200, res)
}
