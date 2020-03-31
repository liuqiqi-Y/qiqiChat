package model

import (
	"qiqiChat/util"
	"time"
)

type ReceiveDetail struct {
	GroupName   string
	StaffName   string
	ProductName string
	Count       int
	Time        time.Time
	Begin       string
	End         string
}

func ReceiveProductDetail(index, size, chararcter int, group_name, staff_name, begin, end string) ([]ReceiveDetail, error) {
	var details []ReceiveDetail
	//var sql string
	sqlStr1 := "SELECT g.`name`, p.`name`, SUM(r.`count`), MAX(r.`time`) " +
		"FROM `fact_record` r JOIN `staff` s ON r.`staff_id` = " +
		"s.`id` JOIN `product` p ON r.`product_id` = " +
		"p.`id` JOIN `group` g ON s.`group_id` = g.`id` "
	if (begin == "" && end == "") && group_name == "" && staff_name == "" {

	} else {
		sqlStr1 = sqlStr1 + " " + "WHERE" + " "
	}
	sqlStr2 := "g.`name` = " + "'" + group_name + "'"
	sqlStr3 := "s.`name` = " + "'" + staff_name + "'"
	sqlStr4 := "r.`time` BETWEEN " + "'" + begin + "'" + " AND " + "'" + end + "'"
	sqlStr5 := " GROUP BY g.`name`, p.`name` HAVING " +
		"p.`name` IN (SELECT `name` FROM `product` WHERE `characteristic` = ?)"
	sqlStr6 := "LIMIT ? OFFSET ?"
	if group_name != "" {
		sqlStr1 = sqlStr1 + sqlStr2
		if staff_name != "" || (begin != "" && end != "") {
			sqlStr1 = sqlStr1 + " AND "
		}
	}
	if staff_name != "" {
		sqlStr1 = sqlStr1 + sqlStr3
		if begin != "" && end != "" {
			sqlStr1 = sqlStr1 + " AND "
		}
	}
	if begin != "" && end != "" {
		sqlStr1 = sqlStr1 + sqlStr4
	}
	sqlStr1 = sqlStr1 + sqlStr5
	sqlStr1 = sqlStr1 + sqlStr6

	//fmt.Println("===>" + sqlStr1)
	rows, err := DB.Query(sqlStr1, chararcter, size, size*(index-1))
	// rows, err := DB.Query("SELECT g.`name`, p.`name`, SUM(r.`count`), MAX(r.time) "+
	// 	"FROM `fact_record` r JOIN `staff` s ON r.`staff_id` = "+
	// 	"s.`id` JOIN `product` p ON r.`product_id` = "+
	// 	"p.`id` JOIN `group` g ON s.`group_id` = g.`id` "+
	// 	"GROUP BY g.`name`, p.`name` HAVING "+
	// 	"p.`name` IN (SELECT `name` FROM `product` WHERE `characteristic` = ?)LIMIT ? OFFSET ?", chararcter, size, size*(index-1))
	if err != nil {
		util.Err.Printf("failed to query: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()
	var detail ReceiveDetail
	for rows.Next() {
		err := rows.Scan(&detail.GroupName, &detail.ProductName, &detail.Count, &detail.Time)
		if err != nil {
			util.Err.Printf("failed to query: %s\n", err.Error())
			return nil, err
		}
		if staff_name != "" {
			detail.StaffName = staff_name
		}
		details = append(details, detail)
	}
	return details, nil
}
func CheckDetailCount(character int, group_name, staff_name, begin, end string) int {
	count := 0
	// _ = DB.QueryRow("SELECT COUNT(*) FROM "+
	// 	"(SELECT p.`name` FROM `fact_record` r JOIN "+
	// 	"`staff` s ON r.`staff_id` = s.`id` JOIN "+
	// 	"`product` p ON r.`product_id` = p.`id` JOIN"+
	// 	"`group` g ON s.`group_id` = g.`id` "+
	// 	"GROUP BY g.`name`, p.`name` HAVING "+
	// 	"p.`name` IN "+
	// 	"(SELECT `name` FROM `product` WHERE `characteristic` = ?))t", character).Scan(&count)

	sqlStr1 := "SELECT p.`name` " +
		"FROM `fact_record` r JOIN `staff` s ON r.`staff_id` = " +
		"s.`id` JOIN `product` p ON r.`product_id` = " +
		"p.`id` JOIN `group` g ON s.`group_id` = g.`id` "
	if (begin == "" && end == "") && group_name == "" && staff_name == "" {

	} else {
		sqlStr1 = sqlStr1 + " " + "WHERE" + " "
	}
	sqlStr2 := "g.`name` = " + "'" + group_name + "'"
	sqlStr3 := "s.`name` = " + "'" + staff_name + "'"
	sqlStr4 := "r.`time` BETWEEN " + "'" + begin + "'" + " AND " + "'" + end + "'"
	sqlStr5 := " GROUP BY g.`name`, p.`name` HAVING " +
		"p.`name` IN (SELECT `name` FROM `product` WHERE `characteristic` = ?)"
	if group_name != "" {
		sqlStr1 = sqlStr1 + sqlStr2
		if staff_name != "" || (begin != "" && end != "") {
			sqlStr1 = sqlStr1 + " AND "
		}
	}
	if staff_name != "" {
		sqlStr1 = sqlStr1 + sqlStr3
		if begin != "" && end != "" {
			sqlStr1 = sqlStr1 + " AND "
		}
	}
	if begin != "" && end != "" {
		sqlStr1 = sqlStr1 + sqlStr4
	}
	sqlStr1 = sqlStr1 + sqlStr5

	_ = DB.QueryRow("SELECT COUNT(*) FROM "+"("+sqlStr1+")"+"t", character).Scan(&count)
	return count
}

type Record struct {
	ID        uint
	StaffID   uint
	ProductID uint
	Count     int
	Type      int
	Time      time.Time
}

func SetRecord(staffID, productID uint, count int, time string) (Record, error) {

	f := 0
	if count >= 0 {
		f = 1 //0表示归还1表示借出
	}
	if count == 0 {
		return Record{}, nil
	}

	tx, _ := DB.Begin()
	result, err := tx.Exec("INSERT INTO `record`(`staff_id`, `product_id`, `count`, `type`, `time`) VALUES(?,?,?,?,?)",
		staffID, productID, count, f, time)
	if err != nil {
		_ = tx.Rollback()
		util.Err.Printf("faile to insert: %s\n", err.Error())
		return Record{}, err
	}
	affect, _ := result.RowsAffected()
	if affect == 0 {
		util.Err.Printf("faile to insert: %s\n", err.Error())
		return Record{}, nil
	}
	exist := CheckFactRecord(staffID, productID)
	if exist == true {
		result, err := tx.Exec("UPDATE `fact_record` SET `count` = `count` + ?, `time` = ? WHERE `staff_id` = ? AND `product_id` = ?",
			count, time, staffID, productID)
		if err != nil {
			_ = tx.Rollback()
			util.Err.Printf("faile to insert: %s\n", err.Error())
			return Record{}, err
		}
		affect, _ := result.RowsAffected()
		if affect == 0 {
			util.Err.Printf("faile to insert: %s\n", err.Error())
			return Record{}, nil
		}
	} else {
		result, err = tx.Exec("INSERT INTO `fact_record`(`staff_id`, `product_id`, `count`, `time`) VALUES(?,?,?,?)", staffID, productID, count, time)
		if err != nil {
			_ = tx.Rollback()
			util.Err.Printf("faile to insert: %s\n", err.Error())
			return Record{}, err
		}
		affect, _ = result.RowsAffected()
		if affect == 0 {
			util.Err.Printf("faile to insert: %s\n", err.Error())
			return Record{}, nil
		}
	}
	_, err = tx.Exec("UPDATE `product` SET `quantity` = `quantity` - ?, `used` = `used` + ? WHERE `id` = ?", count, count, productID)
	if err != nil {
		_ = tx.Rollback()
		util.Err.Printf("faile to insert: %s\n", err.Error())
		return Record{}, err
	}
	record := Record{
		StaffID:   staffID,
		ProductID: productID,
		Count:     count,
		Type:      f,
	}
	_ = tx.QueryRow("SELECT `id`, `time` "+
		"FROM `record` WHERE `staff_id` = ? AND `product_id` = ? AND `count` = ? AND `type` = ? AND `time` = ?",
		staffID, productID, count, f, time).Scan(&record.ID, &record.Time)
	_ = tx.Commit()
	return record, nil
}
func CheckFactRecord(staff_id, product_id uint) bool {
	count := 0
	_ = DB.QueryRow("SELECT COUNT(*) FROM `fact_record` WHERE `staff_id` = ? AND `product_id` = ?", staff_id, product_id).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}
func CheckProductID(productName string, character int) uint {
	id := 0
	_ = DB.QueryRow("SELECT `id` FROM `product` WHERE `name` = ? AND `characteristic` = ? AND `status` = 1", productName, character).Scan(&id)
	return uint(id)
}
func CheckStaffID(staffName, groupName string) uint {
	groupID := 0
	_ = DB.QueryRow("SELECT `id` FROM `group` WHERE `name` = ?", groupName).Scan(&groupID)
	if groupID == 0 {
		return 0
	}
	staffID := 0
	_ = DB.QueryRow("SELECT `id` FROM `staff` WHERE `name` = ? AND `group_id` = ? AND `status` = 1 LIMIT 1", staffName, uint(groupID)).Scan(&staffID)
	return uint(staffID)
}

type RecordData struct {
	StaffID   uint
	ProductID uint
	Count     int
	Time      string
}

func SetRecord1(rd []RecordData) ([]Record, error) {

	var arr []Record
	tx, _ := DB.Begin()
	for _, v := range rd {
		f := 0
		if v.Count >= 0 {
			f = 1 //0表示归还1表示借出
		}
		if v.Count == 0 {
			continue
		}
		result, err := tx.Exec("INSERT INTO `record`(`staff_id`, `product_id`, `count`, `type`, `time`) VALUES(?,?,?,?,?)",
			v.StaffID, v.ProductID, v.Count, f, v.Time)
		if err != nil {
			_ = tx.Rollback()
			util.Err.Printf("faile to insert: %s\n", err.Error())
			return []Record{}, err
		}
		affect, _ := result.RowsAffected()
		if affect == 0 {
			util.Err.Printf("faile to insert: %s\n", err.Error())
			return []Record{}, nil
		}
		exist := CheckFactRecord(v.StaffID, v.ProductID)
		if exist == true {
			result, err := tx.Exec("UPDATE `fact_record` SET `count` = `count` + ?, `time` = ? WHERE `staff_id` = ? AND `product_id` = ?",
				v.Count, v.Time, v.StaffID, v.ProductID)
			if err != nil {
				_ = tx.Rollback()
				util.Err.Printf("faile to insert: %s\n", err.Error())
				return []Record{}, err
			}
			affect, _ := result.RowsAffected()
			if affect == 0 {
				util.Err.Printf("faile to insert: %s\n", err.Error())
				return []Record{}, nil
			}
		} else {
			result, err = tx.Exec("INSERT INTO `fact_record`(`staff_id`, `product_id`, `count`, `time`) VALUES(?,?,?,?)", v.StaffID, v.ProductID, v.Count, v.Time)
			if err != nil {
				_ = tx.Rollback()
				util.Err.Printf("faile to insert: %s\n", err.Error())
				return []Record{}, err
			}
			affect, _ = result.RowsAffected()
			if affect == 0 {
				util.Err.Printf("faile to insert: %s\n", err.Error())
				return []Record{}, nil
			}
		}
		_, err = tx.Exec("UPDATE `product` SET `quantity` = `quantity` - ?, `used` = `used` + ? WHERE `id` = ?", v.Count, v.Count, v.ProductID)
		if err != nil {
			_ = tx.Rollback()
			util.Err.Printf("faile to insert: %s\n", err.Error())
			return []Record{}, err
		}
		record := Record{
			StaffID:   v.StaffID,
			ProductID: v.ProductID,
			Count:     v.Count,
			Type:      f,
		}
		_ = tx.QueryRow("SELECT `id`, `time` "+
			"FROM `record` WHERE `staff_id` = ? AND `product_id` = ? AND `count` = ? AND `type` = ? AND `time` = ?",
			v.StaffID, v.ProductID, v.Count, f, v.Time).Scan(&record.ID, &record.Time)
		arr = append(arr, record)
	}
	_ = tx.Commit()
	return arr, nil
}
func GetOneGroupOneProductRecord(staffID []uint, productID uint) int {
	sum := 0
	for _, v := range staffID {
		count := 0
		_ = DB.QueryRow("SELECT `count` FROM `fact_record` WHERE `staff_id` = ? AND `product_id` = ?", v, productID).Scan(&count)
		sum += count
	}
	return sum
}
