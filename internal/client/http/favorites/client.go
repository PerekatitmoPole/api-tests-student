package favorites

import (
	"net/http"
	"testing"

	httpclient "api-tests-template/internal/client/http"
)

const (
	favoritesPath               = "/favorites"
	favoriteByAdvertisementPath = "/favorites/%s"
	favoritesBulkPath           = "/favorites/bulk"
	maxFavoritesPageSize        = "50"
)

// HttpGetFavorites запрашивает избранное одной страницей максимального размера.
func HttpGetFavorites(t *testing.T, token string) *http.Response {
	t.Helper()
	return httpclient.NewRequest(token).Get(favoritesPath).
		Query("page_size", maxFavoritesPageSize).Expect(t).End().Response
}

// HttpPostFavorite добавляет объявление в избранное покупателя.
func HttpPostFavorite(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()
	request := struct {
		AdvertisementID string `json:"advertisement_id,omitempty"`
	}{AdvertisementID: advertisementID}
	return httpclient.NewRequest(token).Post(favoritesPath).
		JSON(request).Expect(t).End().Response
}

// HttpDeleteFavorite удаляет одно объявление из избранного.
func HttpDeleteFavorite(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()
	return httpclient.NewRequest(token).
		Deletef(favoriteByAdvertisementPath, advertisementID).Expect(t).End().Response
}

// HttpPostFavoritesBulk добавляет список объявлений для подготовки сценария лимита.
func HttpPostFavoritesBulk(t *testing.T, token string, advertisementIDs []string) *http.Response {
	t.Helper()
	return httpclient.NewRequest(token).Post(favoritesBulkPath).
		JSON(map[string][]string{"advertisement_ids": advertisementIDs}).Expect(t).End().Response
}
