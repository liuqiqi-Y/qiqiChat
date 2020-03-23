package service

import (
	"qiqiChat/model"
	"qiqiChat/serializer"
)

type Staff struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Number  string `form:"number" json:"number" binding:"required,min=7,max=7"`
	Email   string `form:"email" json:"email" binding:"email"`
	GroupID uint   `form:"group_id" json:"group_id" binding:"required"`
}

func (s *Staff) InsertStaff() serializer.Response {
	exist := model.CheckStaff(s.Number)
	if exist == true {
		return serializer.Err(40004, "工号重复，无法录入系统。", nil)
	}
	exist = model.CheckGroupByID(s.GroupID)
	if exist == false {
		return serializer.Err(40003, "没有该分组", nil)
	}
	staff, err := model.InsertStaff(s.Name, s.Number, s.Email, s.GroupID)
	if err != nil {
		return serializer.DBErr("", nil)
	}
	return serializer.StaffResponse(staff)
}

func DelStaff(id uint) serializer.Response {
	exist := model.CheckStaffByID(id)
	if exist == false {
		return serializer.Err(40003, "没有该员工", nil)
	}
	success := model.DelStaff(id)
	if success == false {
		return serializer.DBErr("", nil)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "删除员工成功",
	}
}
