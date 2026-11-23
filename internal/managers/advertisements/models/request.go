package models

const (
	TitleField       = "title"
	DescriptionField = "description"
	PriceField       = "price"
	QuantityField    = "quantity"
)

type CreateAdvertisementRequest struct {
	Fields map[string]string
	Photos [][]byte
}
