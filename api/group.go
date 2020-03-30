package api

import (
	"qiqiChat/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GroupInfo(c *gin.Context) {
	var service service.GroupInfo
	res := service.GetGroups()
	c.JSON(200, res)
}
func GroupAdd(c *gin.Context) {
	var service service.GroupAdd
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddGroup()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func DelGroup(c *gin.Context) {
	var service service.GroupInfo
	s := c.Param("GroupID")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
		return
	}
	res := service.DelGroup(uint(id))
	c.JSON(200, res)
}
