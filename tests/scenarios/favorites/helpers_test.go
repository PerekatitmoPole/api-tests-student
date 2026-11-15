package favorites

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/tests/fixtures"
	"api-tests-template/tests/internal/manager"
)

func newAdvertisementID() string {
	var id [16]byte
	rand.Read(id[:])
	id[6] = (id[6] & 0x0f) | 0x40
	id[8] = (id[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", id[:4], id[4:6], id[6:8], id[8:10], id[10:])
}

func cleanupFavorite(t *testing.T, token, advertisementID string) {
	t.Helper()
	t.Cleanup(func() {
		t.Logf("Очистка: удаляем объявление %s из избранного", advertisementID)
		manager.DeleteFavorite(t, token, advertisementID).RequireCleanupStatus(t)
	})
}

func addFavoriteForSetup(t *testing.T, token, advertisementID string) {
	t.Helper()
	cleanupFavorite(t, token, advertisementID)
	t.Logf("Шаг: добавляем объявление %s в избранное", advertisementID)
	manager.AddFavorite(t, token, advertisementID).RequireStatus(t, http.StatusCreated)
}

func listFavorites(t *testing.T, token string) manager.FavoritesList {
	t.Helper()
	result := manager.GetFavorites(t, token)
	result.RequireStatus(t, http.StatusOK)
	list := manager.Decode[manager.FavoritesList](t, result)
	require.NotNil(t, list.Total, "список избранного не содержит поле total")
	require.Len(t, list.Items, *list.Total, "ответ должен содержать все записи избранного")
	return list
}

func requireFavoriteState(t *testing.T, token string, ad *fixtures.PublishedAdvertisement, expectedPresent bool) {
	t.Helper()
	t.Logf("Шаг: проверяем наличие %s в GET /favorites, ожидаем %t", ad.ID, expectedPresent)
	list := listFavorites(t, token)
	found := slices.ContainsFunc(list.Items, func(item manager.Favorite) bool {
		return item.Advertisement.ID == ad.ID
	})
	require.Equal(t, expectedPresent, found, "состояние избранного для объявления %s (%q)", ad.ID, ad.Title)
	t.Logf("Наличие объявления в избранном подтверждено: %t", found)
}
