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
