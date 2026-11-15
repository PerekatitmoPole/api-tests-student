package manager

import (
	"encoding/json"
	"net/http"
	"testing"

	catalog "api-tests-template/internal/client/http/advertisements"
	"api-tests-template/tests/internal/client"
)

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

type Advertisement struct {
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

type SearchItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func SearchAdvertisements(t *testing.T, token, title string) []SearchItem {
	t.Helper()
	result := readResponse(t, catalog.HttpGetAdvertisements(t, token, title))
	return Decode[struct {
		Items []SearchItem `json:"items"`
	}](t, result).Items
}

func CreateAdvertisement(t *testing.T, token string, request CreateAdvertisementRequest) Result {
	t.Helper()
	return readResponse(t, client.HttpPostAdvertisement(t, token, request.Fields, request.Photos))
}

func GetAdvertisement(t *testing.T, token, id string) Advertisement {
	t.Helper()
	result := readResponse(t, client.HttpGetAdvertisement(t, token, id))
	result.RequireStatus(t, http.StatusOK)
	return Decode[Advertisement](t, result)
}

func DeleteAdvertisement(t *testing.T, token, id string) Result {
	t.Helper()
	return readResponse(t, client.HttpDeleteAdvertisement(t, token, id))
}

func GetAdvertisementPhotos(t *testing.T, token, id string) []Photo {
	t.Helper()
	result := readResponse(t, client.HttpGetAdvertisementPhotos(t, token, id))
	result.RequireStatus(t, http.StatusOK)
	return Decode[[]Photo](t, result)
}
