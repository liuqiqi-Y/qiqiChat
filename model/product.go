package model

import (
	"qiqiChat/util"
	"time"
)

type Product struct {
	ID             uint
	Name           string
	Characteristic int //产品分类 1固定资产 0低值易耗
	Quantity       int //当前库存量
	Used           int //已消耗数量
	Created_at     time.Time
	Updated_at     time.Time
	Status         int
	//Count          int
}

func GetProducts(index, size, chararcter int) ([]Product, error) {
	var products []Product
	rows, err := DB.Query("SELECT `id`, `name`, `characteristic`, `quantity`, `used`, `created_at`, `status` FROM `product` WHERE `characteristic` = ? AND `status` = 1 LIMIT ? OFFSET ?", chararcter, size, size*(index-1))
	if err != nil {
		util.Err.Printf("failed to query: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()
	var product Product
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Characteristic, &product.Quantity, &product.Used, &product.Created_at, &product.Status)
		if err != nil {
			util.Err.Printf("failed to query: %s\n", err.Error())
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
func GetProductByName(name string, character int) (Product, error) {
	var product Product
	err := DB.QueryRow("SELECT `id`, `name`, `characteristic`, `quantity`, `used`, `created_at`, `status` FROM `product` WHERE `status` = 1 AND `characteristic` = ?  AND `name` = ?", character, name).Scan(
		&product.ID, &product.Name, &product.Characteristic, &product.Quantity, &product.Used, &product.Created_at, &product.Status)
	if err != nil {
		util.Err.Printf("failed to query: %s\n", err.Error())
		return Product{}, err
	}
	return product, nil
}
func CheckProductByName(name string, character int) bool {
	count := 0
	_ = DB.QueryRow("SELECT COUNT(*) FROM `product` WHERE `name` = ? AND `characteristic` = ? AND `status` = 1", name, character).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}

func GetProductsByTime(index, size, character int, begin, end string) ([]Product, error) {
	if begin == "" || end == "" {
		return GetProducts(index, size, character)
	}
	rows, err := DB.Query("SELECT `id`,`name`,`characteristic`,`quantity`,`used`,`created_at`,`status` FROM `product` WHERE `status` = 1 AND `characteristic` = ? AND `id` IN ( SELECT `product_id` FROM `record` WHERE `type` = 1 AND `time` BETWEEN ? AND ? GROUP BY `product_id` ) LIMIT ? OFFSET ?", character, begin, end, size, size*(index-1))
	if err != nil {
		util.Err.Printf("failed to query: %s\n", err.Error())
		return nil, err
	}
	defer rows.Close()
	var product Product
	var products []Product
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Characteristic, &product.Quantity, &product.Used, &product.Created_at, &product.Status)
		if err != nil {
			util.Err.Printf("failed to query: %s\n", err.Error())
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
func ModifyProductCount(id uint, count int) bool {
	result, err := DB.Exec("UPDATE `product` SET `quantity` = `quantity` + ? WHERE `id` = ?", count, id)
	if err != nil {
		util.Err.Printf("failed to update: %s\n", err.Error())
		return false
	}
	affect, _ := result.RowsAffected()
	if affect == 0 {
		util.Err.Println("failed to update")
		return false
	}
	return true
}
func CheckProductByID(id uint) bool {
	count := 0
	_ = DB.QueryRow("SELECT COUNT(*) FROM `product` WHERE `id` = ? AND `status` = 1", id).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}
func AddProduct(name string, quantity int, character int) (Product, error) {
	tx, _ := DB.Begin()
	stmt, err := tx.Prepare("INSERT INTO `product`(`name`, `characteristic`, `quantity`, `status`) VALUES(?,?,?,?)")
	if err != nil {
		_ = tx.Rollback()
		util.Err.Println("Failed to prepare sql statement: ", err.Error())
		return Product{}, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, character, quantity, 1)
	if err != nil {
		_ = tx.Rollback()
		util.Err.Println("Failed to insert a product: ", err.Error())
		return Product{}, err
	}
	p := Product{}
	_ = tx.QueryRow("SELECT `id`, `name`, `characteristic`, `quantity`, `used`, `created_at`, `status` FROM `product` WHERE `name` = ? AND `characteristic` = ?", name, character).Scan(
		&p.ID, &p.Name, &p.Characteristic, &p.Quantity, &p.Used, &p.Created_at, &p.Status)
	tx.Commit()
	return p, nil
}
func DelProduct(name string, character int) bool {
	result, err := DB.Exec("UPDATE `product` SET `status` = 0 WHERE `name` = ? AND `characteristic` = ?", name, character)
	if err != nil {
		util.Err.Printf("failed to update: %s\n", err.Error())
		return false
	}
	affect, _ := result.RowsAffected()
	if affect == 0 {
		util.Err.Println("failed to update")
		return false
	}
	return true
}

func CheckProductBytime(begin string, end string) bool {
	count := 0
	_ = DB.QueryRow("SELECT COUNT(*) FROM `record` WHERE `type` = 1 AND `time` BETWEEN ? AND ?", begin, end).Scan(&count)
	if count > 0 {
		return true
	}
	return false
}
func CheckProductCount(character int) int {
	count := 0
	_ = DB.QueryRow("SELECT COUNT(*) FROM `product` WHERE `status` = 1 AND `characteristic` = ?", character).Scan(&count)
	return count
}

func ModifyProductName(oldName string, newName string, character int) (Product, error) {
	var product Product
	tx, _ := DB.Begin()
	result, err := tx.Exec("UPDATE `product` SET `name` = ? WHERE `name` = ? AND `characteristic` = ? AND `status` = 1", newName, oldName, character)
	if err != nil {
		_ = tx.Rollback()
		util.Err.Printf("failed to update: %s\n", err.Error())
		return Product{}, err
	}
	affect, _ := result.RowsAffected()
	if affect == 0 {
		_ = tx.Rollback()
		util.Err.Printf("failed to update: %s\n", err.Error())
		return Product{}, err
	}
	_ = tx.QueryRow("SELECT `id`, `name`, `characteristic`, `quantity`, `used`, `created_at`, `status` FROM `product` WHERE `status` = 1 AND `characteristic` = ?  AND `name` = ?", character, newName).Scan(
		&product.ID, &product.Name, &product.Characteristic, &product.Quantity, &product.Used, &product.Created_at, &product.Status)
	tx.Commit()
	return product, nil
}
