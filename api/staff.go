package api

import (
	"qiqiChat/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddStaff(c *gin.Context) {
	var service service.Staff
	if err := c.ShouldBind(&service); err == nil {
		res := service.InsertStaff()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func DelStaff(c *gin.Context) {
	s := c.Param("StaffID")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
	}
	res := service.DelStaff(uint(id))
	c.JSON(200, res)
}
func UpdateStaff(c *gin.Context) {
	var service service.StaffUpdate
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpdateStaff()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
func GetStaffes(c *gin.Context) {
	s := c.Param("GroupID")
	id, err := strconv.Atoi(s)
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
	}
	res := service.GetStaffes(uint(id))
	c.JSON(200, res)
}
