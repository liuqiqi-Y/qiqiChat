package model

import (
	"qiqiChat/util"
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
