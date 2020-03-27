package serializer

import "qiqiChat/model"

type Staff struct {
	ID        uint   `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Number    string `json:"number,omitempty"`
	Email     string `json:"email,omitempty"`
	GroupID   uint   `json:"group_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Status    int    `json:"status,omitempty"`
}

func buildStaff(staff model.Staff) Staff {
	return Staff{
		ID:        staff.ID,
		Name:      staff.Name,
		Number:    staff.Number,
		Email:     staff.Email,
		GroupID:   staff.GroupID,
		CreatedAt: staff.CreatedAt.Format("2006-01-02 15:04:05"),
		Status:    staff.Status,
	}
}
func StaffesResponse(staffes []model.Staff) Response {
	var s []Staff
	for _, v := range staffes {
		s = append(s, buildStaff(v))
	}
	return Response{
		Data: s,
	}
}
func StaffResponse(staff model.Staff) Response {
	return Response{
		Data: buildStaff(staff),
		Msg:  "员工信息修改成共",
	}
}
