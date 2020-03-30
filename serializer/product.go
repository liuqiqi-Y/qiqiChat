package serializer

import "qiqiChat/model"

type Product struct {
	ID             uint   `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Characteristic int    `json:"characteristic"`
	Quantity       int    `json:"quantity,omitempty"`
	Used           int    `json:"used"`
	CreatedAt      string `json:"created_at,omitempty"`
	Status         int    `json:"status"`
}
type ProductEmpty struct {
	ID             uint   `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Characteristic int    `json:"characteristic,omitempty"`
	Quantity       int    `json:"quantity,omitempty"`
	Used           int    `json:"used,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	Status         int    `json:"status,omitempty"`
}
type ProductList struct {
	Total     int       `json:"total"`
	TotalPage int       `json:"total_page"`
	List      []Product `json:"list"`
}
type ProductListEmpty struct {
	Total     int            `json:"total"`
	TotalPage int            `json:"total_page"`
	List      []ProductEmpty `json:"list"`
}

func buildProduct(p model.Product) Product {
	return Product{
		ID:             p.ID,
		Name:           p.Name,
		Characteristic: p.Characteristic,
		Quantity:       p.Quantity,
		Used:           p.Used,
		CreatedAt:      p.Created_at.Format("2006-01-02 15:04:05"),
		Status:         p.Status,
	}
}
func ProductsResponse(products []model.Product, total int, size int) Response {

	var s []Product
	for _, v := range products {
		s = append(s, buildProduct(v))
	}
	var pl ProductList
	pl.List = s
	pl.Total = total
	pl.TotalPage = (total + (total % size)) / size
	if total+(total%size) < size {
		pl.TotalPage = 1
	} else {
		pl.TotalPage = (total + (total % size)) / size
	}
	return Response{
		Data: pl,
	}
}
func ProductResponse(product model.Product) Response {
	return Response{
		Data: buildProduct(product),
	}
}
