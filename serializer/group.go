package serializer

import "qiqiChat/model"

// Group 分组序列化器
type Group struct {
	ID        uint   `json:"id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	Status    int    `json:"status,omitempty"`
}

// BuildGroups 序列化分组
func BuildGroups(groups []model.Group) []Group {
	var groupArr []Group
	for _, v := range groups {
		g := Group{
			ID:        v.ID,
			GroupName: v.Name,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			Status:    v.Status,
		}
		groupArr = append(groupArr, g)
	}
	return groupArr
}

// BuildGroupsResponse 返回分组响应
func BuildGroupsResponse(groups []model.Group) Response {
	return Response{
		Data: BuildGroups(groups),
	}
}
func BuildGroup(group model.Group) Group {
	return Group{
		ID:        group.ID,
		GroupName: group.Name,
		CreatedAt: group.CreatedAt.Format("2006-01-02 15:04:05"),
		Status:    group.Status,
	}
}
func BuildGroupResponse(group model.Group) Response {
	return Response{
		Data: BuildGroup(group),
		Msg:  "新增组别成功",
	}
}
