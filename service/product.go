package service

import (
	"qiqiChat/model"
	"qiqiChat/serializer"
)

type ProductInfo struct {
	Index     int
	Size      int
	Character int
}

func (p *ProductInfo) GetProducts() serializer.Response {
	if p.Index <= 0 || p.Size <= 0 || (p.Character != 0 && p.Character != 1) {
		return serializer.ParamErr("", nil)
	}
	count := model.CheckProductCount(p.Character)
	if count <= 0 {
		return serializer.Response{
			Code: 0,
			Data: serializer.ProductList{},
			Msg:  "没有物品库存",
		}
	}
	products, err := model.GetProducts(p.Index, p.Size, p.Character)
	if err != nil {
		return serializer.DBErr("", nil)
	}
	return serializer.ProductsResponse(products, count, p.Size)
}

type ProductByName struct {
	Name      string `form:"name" json:"name" binding:"required"`
	Character int    //`path:"characteristic"`
}

func (p *ProductByName) GetProductByName() serializer.Response {
	if p.Character != 0 && p.Character != 1 {
		return serializer.ParamErr("", nil)
	}
	exist := model.CheckProductByName(p.Name, p.Character)
	if exist == false {
		return serializer.Response{
			Code: 0,
			Data: serializer.ProductEmpty{},
			Msg:  "没有该物品",
		}
	}
	product, err := model.GetProductByName(p.Name, p.Character)
	if err != nil {
		return serializer.DBErr("", nil)
	}
	return serializer.ProductResponse(product)
}

type ProductsByTime struct {
	Index     int `form:"index" json:"index" binding:"required"`
	Size      int `form:"size" json:"size" binding:"required"`
	Character int
	Begin     string `form:"begin" json:"begin" time_format:"2006-01-02"`
	End       string `form:"end" json:"end" time_format:"2006-01-02"`
}

func (p *ProductsByTime) GetProductsByTime() serializer.Response {
	if p.Index <= 0 || p.Size <= 0 || (p.Character != 0 && p.Character != 1) {
		return serializer.ParamErr("", nil)
	}

	if p.Begin != "" && p.End != "" {
		exist := model.CheckProductBytime(p.Begin, p.End)
		if exist == false {
			return serializer.Response{
				Code: 0,
				Data: []model.Product{},
				Msg:  "该时间段没有物品消耗",
			}
		}
	}
	products, err := model.GetProductsByTime(p.Index, p.Size, p.Character, p.Begin, p.End)
	if err != nil {
		return serializer.DBErr("", nil)
	}
	return serializer.ProductsResponse(products, 0, 0)
}

type ProductCount struct {
	ID    uint `form:"id" json:"id" binding:"required"`
	Count int  `form:"count" json:"count" binding:"required"`
}

func (p *ProductCount) ModifyProductCount() serializer.Response {
	exist := model.CheckProductByID(p.ID)
	if exist == false {
		return serializer.Err(40003, "没有该物品", nil)
	}
	success := model.ModifyProductCount(p.ID, p.Count)
	if success == false {
		return serializer.DBErr("更新失败", nil)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "修改成功",
	}
}

type ProductAdd struct {
	Name      string `form:"name" json:"name" binding:"required"`
	Character int
	Quantity  int `form:"quantity" json:"quantity" binding:"required"`
}

func (p *ProductAdd) AddProduct() serializer.Response {
	exist := model.CheckProductByName(p.Name, p.Character)
	if exist == true {
		return serializer.Err(40002, "该物品已存在", nil)
	}
	product, err := model.AddProduct(p.Name, p.Quantity, p.Character)
	if err != nil {
		return serializer.DBErr("", nil)
	}
	return serializer.ProductResponse(product)
}
func (p *ProductByName) DelProduct() serializer.Response {
	if p.Character != 0 && p.Character != 1 {
		return serializer.ParamErr("", nil)
	}
	exist := model.CheckProductByName(p.Name, p.Character)
	if exist == false {
		return serializer.Err(40003, "该物品并不存在", nil)
	}
	success := model.DelProduct(p.Name, p.Character)
	if success == false {
		return serializer.DBErr("删除失败", nil)
	}
	return serializer.Response{
		Code: 0,
		Msg:  "删除成功",
	}
}

type ProductName struct {
	OldName   string `form:"old_name" binding:"required"`
	NewName   string `form:"new_name" binding:"required"`
	Character int
}

func (p *ProductName) ModifyProductName() serializer.Response {
	if p.Character != 0 && p.Character != 1 {
		return serializer.ParamErr("", nil)
	}
	exist := model.CheckProductByName(p.OldName, p.Character)
	if exist == false {
		return serializer.Err(40003, "无此类别，请重新输入。", nil)
	}
	product, err := model.ModifyProductName(p.OldName, p.NewName, p.Character)
	if err != nil {
		return serializer.DBErr("更新失败", nil)
	}
	return serializer.ProductResponse(product)
}
