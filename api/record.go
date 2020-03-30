package api

import (
	"qiqiChat/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReceiveDetail(c *gin.Context) {
	var p service.RecordInfo
	index := c.Query("index")
	size := c.Query("size")
	character := c.Param("characteristic")
	var err error
	//err = c.ShouldBindQuery(&p)
	p.GroupName = c.Query("group_name")
	p.StaffName = c.Query("staff_name")
	p.Begin = c.Query("begin")
	p.End = c.Query("end")
	if err != nil {
		c.JSON(200, ErrorResponse(nil))
		return
	}
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
	res := p.GetReceiveDetails()
	//s, _ := json.Marshal(&res)
	//fmt.Println("====>" + string(s))
	c.JSON(200, res)
}
