package service

import (
	"qiqiChat/serializer"

	"qiqiChat/model"
)

type GroupInfo struct {
}

func (g *GroupInfo) GetGroups() serializer.Response {
	groups, err := model.GetGroups()
	if len(groups) == 0 {
		return serializer.Response{
			Code: 0,
			Data: []serializer.Group{},
			Msg:  "没有分组",
		}
	}
	if err != nil {
		return serializer.DBErr("", err)
	}
	return serializer.BuildGroupsResponse(groups)
}

type GroupAdd struct {
	GroupName string `form:"group_name" json:"group_name" binding:"required"`
}

func (g *GroupAdd) AddGroup() serializer.Response {
	exist := model.CheckGroupExist(g.GroupName)
	if exist == true {
		return serializer.Err(40002, "组别重复，请重新输入。", nil)
	}
	group, err := model.InsertGroup(g.GroupName)
	if err != nil {
		return serializer.DBErr("", nil)
	}
	return serializer.BuildGroupResponse(group)
}
func (g *GroupInfo) DelGroup(id uint) serializer.Response {
	exist := model.CheckGroupByID(id)
	if exist == false {
		return serializer.Err(40003, "没有该分组", nil)
	}
	success := model.DeleteGroup(id)
	if success == false {
		return serializer.DBErr("", nil)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "删除成功",
	}
}
