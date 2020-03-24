package model

import (
	"fmt"
	"qiqiChat/util"
	"strconv"
	"strings"
	"time"
)

type Staff struct {
	ID        uint
	Name      string
	Number    string
	Email     string
	GroupID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    int
}

func InsertStaff(name, number, email string, groupId uint) (Staff, error) {
	tx, _ := DB.Begin()
	stmt, err := tx.Prepare("INSERT INTO `staff`(`name`, `number`, `email`, group_id, status) VALUES(?,?,?,?,?)")
	if err != nil {
		tx.Rollback()
		util.Err.Println("Failed to prepare sql statement: ", err.Error())
		return Staff{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, number, email, groupId, 1)
	if err != nil {
		tx.Rollback()
		util.Err.Println("Failed to insert a group: ", err.Error())
		return Staff{}, err
	}
	s := Staff{}
	_ = tx.QueryRow("SELECT id, `name`, `number`, `email`, group_id, created_at, status FROM `staff` WHERE `name` = ?", name).Scan(
		&s.ID, &s.Name, &s.Number, &s.Email, &s.GroupID, &s.CreatedAt, &s.Status)
	tx.Commit()
	return s, nil
}
func CheckStaff(number string) bool {
	count := 0
	_ = DB.QueryRow("SELECT COUNT(*) FROM `staff` WHERE `number` = ? AND `status` = 1", number).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}
func CheckStaffByID(id uint) bool {
	count := 0
	_ = DB.QueryRow("SELECT COUNT(*) FROM `staff` WHERE `id` = ? AND `status` = 1", id).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}

func DelStaff(id uint) bool {
	stmt, err := DB.Prepare("UPDATE `staff` SET `status` = 0 WHERE id = ?")
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

type updateField struct {
	name     string
	number   string
	email    string
	group_id uint
}

func (u *updateField) buildUpdateStmt() string {
	var strs []string
	fix := "UPDATE `staff` SET "
	strs = append(strs, fix)
	if u.name != "" {
		strs = append(strs, "`name` = "+u.name+", ")
	}
	if u.number != "" {
		strs = append(strs, "`number` = "+u.number+", ")
	}
	if u.email != "" {
		strs = append(strs, "`email` = "+u.email+", ")
	}
	if u.group_id != 0 {
		s := strconv.Itoa(int(u.group_id))
		strs = append(strs, "`group_id` = "+s+" ")
	}
	strs = append(strs, "WHERE `status` = 1 AND id = ?")
	return strings.Join(strs, "")
}
func UpdateStaff(name, number, email string, group_id, id uint) (Staff, error) {
	u := updateField{
		name:     name,
		number:   number,
		email:    email,
		group_id: group_id,
	}
	str := u.buildUpdateStmt()
	tx, _ := DB.Begin()
	result, err := tx.Exec(str, id)
	if err != nil {
		tx.Rollback()
		util.Err.Printf("failed to update: %s\n", err.Error())
		return Staff{}, err
	}
	affect, _ := result.RowsAffected()
	if affect == 0 {
		_ = tx.Rollback()
		util.Warn.Printf("no row affect\n")
		return Staff{}, fmt.Errorf("no row affect")
	}
	staff := Staff{}
	err = tx.QueryRow("SELECT `id`, `name`, `number`, `email`, `group_id`, `created_at`, `status` FROM `staff` WHERE `status` = 1 AND `id` = ?", id).Scan(
		&staff.ID, &staff.Name, &staff.Number, &staff.Email, &staff.GroupID, &staff.CreatedAt, &staff.Status)
	if err != nil {
		_ = tx.Rollback()
		util.Err.Printf("failed to update: %s\n", err.Error())
		return Staff{}, err
	}
	tx.Commit()
	return staff, nil
}

func GetStaffes(id uint) ([]Staff, error) {
	rows, err := DB.Query("SELECT `id`, `name`, `number`, `email`, `group_id`, `created_at`, `status` FROM `staff` WHERE `status` = 1 AND `group_id` = ?", id)
	if err != nil {
		util.Err.Printf("failed to query: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()
	staff := Staff{}
	staffes := []Staff{}
	for rows.Next() {
		err := rows.Scan(&staff.ID, &staff.Name, &staff.Number, &staff.Email, &staff.GroupID, &staff.CreatedAt, &staff.Status)
		if err != nil {
			util.Err.Printf("failed to get staff info: %s\n", err.Error())
			return nil, err
		}
		staffes = append(staffes, staff)
	}
	return staffes, nil
}
func CheckStaffByGroupID(id uint) bool {
	count := 0
	_ = DB.QueryRow("SELECT COUNT(*) FROM `staff` WHERE `group_id` = ? AND `status` = 1", id).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}
