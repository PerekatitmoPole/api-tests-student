package client

import (
	"net/http"
	"testing"
)

const (
	favoritesPath               = "/favorites"
	favoriteByAdvertisementPath = "/favorites/%s"
	favoritesBulkPath           = "/favorites/bulk"
	maxFavoritesPageSize        = "50"
)

func HttpGetFavorites(t *testing.T, token string) *http.Response {
	t.Helper()
	return newRequest(token).Get(favoritesPath).
		Query("page_size", maxFavoritesPageSize).Expect(t).End().Response
}

func HttpPostFavorite(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()
	request := struct {
		AdvertisementID string `json:"advertisement_id,omitempty"`
	}{AdvertisementID: advertisementID}
	return newRequest(token).Post(favoritesPath).
		JSON(request).Expect(t).End().Response
}

func HttpDeleteFavorite(t *testing.T, token, advertisementID string) *http.Response {
	t.Helper()
	return newRequest(token).
		Deletef(favoriteByAdvertisementPath, advertisementID).Expect(t).End().Response
}

func HttpPostFavoritesBulk(t *testing.T, token string, advertisementIDs []string) *http.Response {
	t.Helper()
	return newRequest(token).Post(favoritesBulkPath).
		JSON(map[string][]string{"advertisement_ids": advertisementIDs}).Expect(t).End().Response
}
