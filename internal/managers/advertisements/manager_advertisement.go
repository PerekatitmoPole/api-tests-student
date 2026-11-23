package advertisements

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
)

// CreateAdvertisement отправляет запрос на публикацию и возвращает ответ API.
func CreateAdvertisement(t *testing.T, token string, request models.CreateAdvertisementRequest) assertions.Result {
	t.Helper()
	return assertions.ReadResponse(t, advertisements.HttpPostAdvertisement(t, token, request.Fields, request.Photos))
}

// GetAdvertisement получает карточку и проверяет успешный ответ.
func GetAdvertisement(t *testing.T, token, id string) models.AdvertisementDetails {
	t.Helper()
	result := assertions.ReadResponse(t, advertisements.HttpGetAdvertisement(t, token, id))
	result.RequireStatus(t, http.StatusOK)
	return assertions.Decode[models.AdvertisementDetails](t, result)
}

// DeleteAdvertisement удаляет объявление и возвращает ответ API.
func DeleteAdvertisement(t *testing.T, token, id string) assertions.Result {
	t.Helper()
	return assertions.ReadResponse(t, advertisements.HttpDeleteAdvertisement(t, token, id))
}

// GetAdvertisementPhotos получает список фотографий объявления.
func GetAdvertisementPhotos(t *testing.T, token, id string) []models.Photo {
	t.Helper()
	result := assertions.ReadResponse(t, advertisements.HttpGetAdvertisementPhotos(t, token, id))
	result.RequireStatus(t, http.StatusOK)
	return assertions.Decode[[]models.Photo](t, result)
}
