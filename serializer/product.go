package serializer

import "qiqiChat/model"

type Product struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Characteristic int    `json:"characteristic"`
	Quantity       int    `json:"quantity"`
	Used           int    `json:"used"`
	CreatedAt      string `json:"created_at"`
	Status         int    `json:"status"`
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
func ProductsResponse(products []model.Product) Response {
	var s []Product
	for _, v := range products {
		s = append(s, buildProduct(v))
	}
	return Response{
		Data: s,
	}
}
func ProductResponse(product model.Product) Response {
	return Response{
		Data: buildProduct(product),
	}
}
