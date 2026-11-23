package models

import "encoding/json"

type AdvertisementsList struct {
	Items []Advertisement `json:"items"`
}

type Advertisement struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type AdvertisementDetails struct {
	ID          string          `json:"id"`
	UserID      string          `json:"user_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Price       json.RawMessage `json:"price"`
	Quantity    json.RawMessage `json:"quantity"`
	Photos      []Photo         `json:"photos"`
}

type Photo struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}
