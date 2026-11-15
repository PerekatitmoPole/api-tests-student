package manager

import (
	"testing"

	"api-tests-template/tests/internal/client"
)

type FavoritesList struct {
	Items []Favorite `json:"items"`
	Total *int       `json:"total"`
}

type Favorite struct {
	Advertisement struct {
		ID string `json:"id"`
	} `json:"advertisement"`
}

func GetFavorites(t *testing.T, token string) Result {
	t.Helper()
	return readResponse(t, client.HttpGetFavorites(t, token))
}

func AddFavorite(t *testing.T, token, advertisementID string) Result {
	t.Helper()
	return readResponse(t, client.HttpPostFavorite(t, token, advertisementID))
}

func DeleteFavorite(t *testing.T, token, advertisementID string) Result {
	t.Helper()
	return readResponse(t, client.HttpDeleteFavorite(t, token, advertisementID))
}

func AddFavoritesBulk(t *testing.T, token string, advertisementIDs []string) Result {
	t.Helper()
	return readResponse(t, client.HttpPostFavoritesBulk(t, token, advertisementIDs))
}
