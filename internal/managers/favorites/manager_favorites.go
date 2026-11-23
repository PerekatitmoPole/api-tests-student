package favorites

import (
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/client/http/favorites"
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

// GetFavorites возвращает ответ сервера на запрос избранного.
func GetFavorites(t *testing.T, token string) assertions.Result {
	t.Helper()
	return assertions.ReadResponse(t, favorites.HttpGetFavorites(t, token))
}

// AddFavorite добавляет объявление и возвращает ответ API.
func AddFavorite(t *testing.T, token, advertisementID string) assertions.Result {
	t.Helper()
	return assertions.ReadResponse(t, favorites.HttpPostFavorite(t, token, advertisementID))
}

// DeleteFavorite удаляет объявление из избранного и возвращает ответ API.
func DeleteFavorite(t *testing.T, token, advertisementID string) assertions.Result {
	t.Helper()
	return assertions.ReadResponse(t, favorites.HttpDeleteFavorite(t, token, advertisementID))
}

// AddFavoritesBulk заполняет избранное для проверки лимита.
func AddFavoritesBulk(t *testing.T, token string, advertisementIDs []string) assertions.Result {
	t.Helper()
	return assertions.ReadResponse(t, favorites.HttpPostFavoritesBulk(t, token, advertisementIDs))
}
