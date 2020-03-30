package service

import (
	"qiqiChat/model"
	"qiqiChat/serializer"
	"regexp"
	"time"
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

type LeadOutTime struct {
	Begin string `form:"begin" json:"begin"`
	End   string `form:"end" json:"end"`
}

func (l *LeadOutTime) LeadingOut() (serializer.Response, string) {
	if (l.Begin == "" && l.End != "") || (l.Begin != "" && l.End == "") {
		return serializer.Err(40001, "开始时间和结束时间不能一个有一个没有", nil), ""
	}
	if l.Begin != "" {
		matched, _ := regexp.MatchString(`((((19|20)\d{2})-(0?[13578]|1[02])-(0?[1-9]|[12]\d|3[01]))|(((19|20)\d{2})-(0?[469]|11)-(0?[1-9]|[12]\d|30))|(((19|20)\d{2})-0?2-(0?[1-9]|1\d|2[0-8]))|((((19|20)([13579][26]|[2468][048]|0[48]))|(2000))-0?2-(0?[1-9]|[12]\d)))$`, l.Begin)
		if matched == false {
			return serializer.Err(40001, "日期时间无效", nil), ""
		}
	}
	if l.End != "" {
		matched, _ := regexp.MatchString(`((((19|20)\d{2})-(0?[13578]|1[02])-(0?[1-9]|[12]\d|3[01]))|(((19|20)\d{2})-(0?[469]|11)-(0?[1-9]|[12]\d|30))|(((19|20)\d{2})-0?2-(0?[1-9]|1\d|2[0-8]))|((((19|20)([13579][26]|[2468][048]|0[48]))|(2000))-0?2-(0?[1-9]|[12]\d)))$`, l.End)
		if matched == false {
			return serializer.Err(40001, "日期时间无效", nil), ""
		}
	}

	result, err := model.GetRecordsByTime(l.Begin, l.End)
	if err != nil {
		return serializer.DBErr("", nil), ""
	}
	if len(result) == 0 {
		return serializer.Response{
			Code: 0,
			Msg:  "没有记录",
		}, ""
	}
	filePath := time.Now().Format("2006-01-02")
	success := model.BuildExcel(result, filePath)
	if success == false {
		return serializer.Response{
			Code: 50002,
			Msg:  "生成excel文件失败",
		}, ""
	}
	return serializer.Response{
		Code: 0,
		Msg:  "导出文件成功",
	}, filePath + "_record.xlsx"
}
