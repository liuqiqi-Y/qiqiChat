package serializer

import "qiqiChat/model"

type RecordList struct {
	Total     int      `json:"total"`
	TotalPage int      `json:"total_page"`
	List      []Detail `json:"list"`
}

type Detail struct {
	GroupName   string `json:"group_name"`
	StaffName   string `json:"staff_name"`
	ProductName string `json:"product_name"`
	Count       int    `json:"count"`
	Time        string `json:"time"`
	Begin       string `json:"begin"`
	End         string `json:"end"`
}

func buildDetail(d model.ReceiveDetail) Detail {
	return Detail{
		GroupName:   d.GroupName,
		StaffName:   d.StaffName,
		ProductName: d.ProductName,
		Count:       d.Count,
		Time:        d.Time.Format("2006-01-02"),
		Begin:       d.Begin,
		End:         d.End,
	}
}

func DetailsResponse(details []model.ReceiveDetail, total int, size int) Response {

	var s []Detail
	for _, v := range details {
		s = append(s, buildDetail(v))
	}
	var pl RecordList
	pl.List = s
	pl.Total = total
	pl.TotalPage = (total + (total % size)) / size
	if total+(total%size) < size {
		pl.TotalPage = 1
	} else {
		pl.TotalPage = (total + (total % size)) / size
	}
	return Response{
		Data: pl,
	}
}