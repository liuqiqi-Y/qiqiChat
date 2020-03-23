package serializer

import "qiqiChat/model"

type Staff struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Number    string `json:"number"`
	Email     string `json:"email"`
	GroupID   uint   `json:"group_id"`
	CreatedAt string `json:"created_at"`
	Status    int    `json:"status"`
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

func StaffResponse(staff model.Staff) Response {
	return Response{
		Data: buildStaff(staff),
	}
}
