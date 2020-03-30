package model

import (
	"qiqiChat/util"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

type Excel struct {
	GroupName      string
	StaffName      string
	StaffNumber    string
	ProductName    string
	Time           time.Time
	Count          int
	Characteristic int
}

func GetRecordsByTime(begin, end string) ([]Excel, error) {
	var excels []Excel

	sql := "SELECT g.`name`, s.`name`, s.`number`, p.`name`, f.`time`, f.`count`, p.`characteristic` FROM " +
		"`fact_record` f JOIN `staff` s ON f.`staff_id` = s.`id` JOIN `product` p ON f.`product_id` = p.`id` JOIN `group` g ON s.`group_id` = g.`id` "
	if begin == "" && end == "" {
		var excel Excel
		rows, err := DB.Query(sql)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&excel.GroupName, &excel.StaffName, &excel.StaffNumber, &excel.ProductName, &excel.Time, &excel.Count, &excel.Characteristic)
			if err != nil {
				util.Err.Printf("failed to query: %s\n", err.Error())
				return nil, err
			}
			excels = append(excels, excel)
		}
		return excels, nil
	}
	sql = sql + "WHERE f.`time` BETWEEN ? AND ?"
	var excel Excel
	rows, err := DB.Query(sql, begin, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&excel.GroupName, &excel.StaffName, &excel.StaffNumber, &excel.ProductName, &excel.Time, &excel.Count, &excel.Characteristic)
		if err != nil {
			util.Err.Printf("failed to query: %s\n", err.Error())
			return nil, err
		}
		excels = append(excels, excel)
	}
	return excels, nil
}

func BuildExcel(excels []Excel, filePath string) bool {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	//var row, row1, row2 *xlsx.Row
	rows := make([]*xlsx.Row, len(excels)+1)
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		util.Err.Printf(err.Error())
		return false
	}
	rows[0] = sheet.AddRow()
	rows[0].SetHeightCM(1)
	cell = rows[0].AddCell()
	cell.Value = "组别"
	cell = rows[0].AddCell()
	cell.Value = "姓名"
	cell = rows[0].AddCell()
	cell.Value = "工号"
	cell = rows[0].AddCell()
	cell.Value = "领用物料"
	cell = rows[0].AddCell()
	cell.Value = "领用时间"
	cell = rows[0].AddCell()
	cell.Value = "领用数量"
	cell = rows[0].AddCell()
	cell.Value = "类别"
	i := 1
	for _, v := range excels {
		rows[i] = sheet.AddRow()
		rows[i].SetHeightCM(1)
		cell = rows[i].AddCell()
		cell.Value = v.GroupName
		cell = rows[i].AddCell()
		cell.Value = v.StaffName
		cell = rows[i].AddCell()
		cell.Value = v.StaffNumber
		cell = rows[i].AddCell()
		cell.Value = v.ProductName
		cell = rows[i].AddCell()
		cell.Value = v.Time.Format("2006-01-02")
		cell = rows[i].AddCell()
		cell.Value = strconv.Itoa(v.Count)
		cell = rows[i].AddCell()
		if v.Characteristic == 0 {
			cell.Value = "易消耗品"
		} else {
			cell.Value = "固定资产"
		}
		i++
	}
	//filePath := time.Now().Format("2006-01-02 15:04:05")
	err = file.Save(filePath + "_record.xlsx")
	if err != nil {
		util.Err.Printf(err.Error())
		return false
	}
	return true
}
