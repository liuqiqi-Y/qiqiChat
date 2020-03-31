package model

import (
	"qiqiChat/util"
	"time"
)

// Group 分组模型
type Group struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    int
}

// GetGroups 获取分组信息
func GetGroups() ([]Group, error) {
	var groups []Group
	rows, err := DB.Query("SELECT id, `name`, created_at, updated_at, `status` FROM `group` WHERE status = 1")
	if err != nil {
		util.Err.Printf("failed to get group info: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()
	var group Group
	for rows.Next() {
		err = rows.Scan(&group.ID, &group.Name, &group.CreatedAt, &group.UpdatedAt, &group.Status)
		if err != nil {
			util.Err.Printf("failed to get group info: %s\n", err.Error())
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
func InsertGroup(name string) (Group, error) {
	tx, _ := DB.Begin()
	stmt, err := tx.Prepare("INSERT INTO `group`(`name`, `status`) VALUES(?,?)")
	if err != nil {
		_ = tx.Rollback()
		util.Err.Println("Failed to prepare sql statement: ", err.Error())
		return Group{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, 1)
	if err != nil {
		_ = tx.Rollback()
		util.Err.Println("Failed to insert a group: ", err.Error())
		return Group{}, err
	}
	g := Group{}
	_ = tx.QueryRow("SELECT id, `name`, created_at, status FROM `group` WHERE `name` = ?", name).Scan(
		&g.ID, &g.Name, &g.CreatedAt, &g.Status)
	tx.Commit()
	return g, nil
}
func CheckGroupExist(name string) bool {
	count := 0
	_ = DB.QueryRow("SELECT COUNT(*) FROM `group` WHERE `name` = ? AND `status` = 1", name).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}
func DeleteGroup(id uint) bool {
	stmt, err := DB.Prepare("UPDATE `group` SET `status` = 0 WHERE id = ?")
	if err != nil {
		util.Err.Printf("failed to prepare sql statement: %s\n", err.Error())
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		util.Err.Printf("failed to update sql: %s\n", err.Error())
		return false
	}
	return true
}
func CheckGroupByID(id uint) bool {
	count := 0
	_ = DB.QueryRow("SELECT COUNT(*) FROM `group` WHERE id = ? AND `status` = 1", id).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}
func GetGroupIDByName(name string) uint {
	count := 0
	_ = DB.QueryRow("SELECT `id` FROM `group` WHERE `name` = ? AND `status` = 1", name).Scan(&count)
	return uint(count)
}
