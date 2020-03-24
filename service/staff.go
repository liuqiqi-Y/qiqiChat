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
		return serializer.Err(40002, "工号重复，无法录入系统。", nil)
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

type StaffUpdate struct {
	ID      uint   `form:"id" json:"id" binding:"required"`
	Name    string `form:"name" json:"name"`
	Number  string `form:"number" json:"number"`
	Email   string `form:"email" json:"email"`
	GroupID uint   `form:"group_id" json:"group_id"`
}

func (s *StaffUpdate) UpdateStaff() serializer.Response {
	exits := model.CheckStaffByID(s.ID)
	if exits == false {
		return serializer.Err(40003, "没有该员工", nil)
	}
	if s.GroupID != 0 {
		exits = model.CheckGroupByID(s.GroupID)
		if exits == false {
			return serializer.Err(40003, "没有该分组", nil)
		}
	}
	if s.Number != "" {
		exits := model.CheckStaff(s.Number)
		if exits == true {
			return serializer.Err(40002, "工号重复，无法录入。", nil)
		}
	}
	staff, err := model.UpdateStaff(s.Name, s.Number, s.Email, s.GroupID, s.ID)
	if err != nil {
		return serializer.DBErr("", nil)
	}
	return serializer.StaffResponse(staff)
}
func GetStaffes(id uint) serializer.Response {
	exist := model.CheckGroupByID(id)
	if exist == false {
		return serializer.Err(40003, "没有该分组", nil)
	}
	exist = model.CheckStaffByGroupID(id)
	if exist == false {
		return serializer.Err(40004, "该分组还没有员工", nil)
	}
	staffes, err := model.GetStaffes(id)
	if err != nil {
		return serializer.DBErr("", nil)
	}
	return serializer.StaffesResponse(staffes)
}
