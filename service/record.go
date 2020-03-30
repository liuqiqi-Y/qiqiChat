package service

import (
	"qiqiChat/model"
	"qiqiChat/serializer"
)

type RecordInfo struct {
	Index     int    `form:"index" json:"index"`
	Size      int    `form:"size" json:"size"`
	GroupName string `form:"group_name json:"group_name"`
	StaffName string `form:"staff_name json:"staff_name:`
	Begin     string `form:"begin json:"begin`
	End       string `form:"end" json:"end"`
	Character int    `form:"character json:"character"`
}

func (r *RecordInfo) GetReceiveDetails() serializer.Response {
	if r.Index <= 0 || r.Size <= 0 || (r.Character != 0 && r.Character != 1) {
		return serializer.ParamErr("", nil)
	}
	count := model.CheckDetailCount(r.Character, r.GroupName, r.StaffName, r.Begin, r.End)
	if count <= 0 {
		return serializer.Response{
			Code: 0,
			Data: serializer.RecordList{
				List: []serializer.Detail{},
			},
			Msg: "该类别没有借用物品记录",
		}
	}
	details, err := model.ReceiveProductDetail(r.Index, r.Size, r.Character, r.GroupName, r.StaffName, r.Begin, r.End)
	if err != nil {
		return serializer.DBErr("", nil)
	}
	if len(details) == 0 {
		return serializer.Response{
			Code: 0,
			Data: serializer.RecordList{
				List: []serializer.Detail{},
			},
			Msg: "",
		}
	}
	return serializer.DetailsResponse(details, count, r.Size)
}
