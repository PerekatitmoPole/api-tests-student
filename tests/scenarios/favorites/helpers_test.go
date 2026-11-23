package favorites

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

// setupFavoriteAdvertisement загружает окружение и готовит покупателя и объявление владельца.
func setupFavoriteAdvertisement(t *testing.T) (string, *fixtures.PublishedAdvertisement) {
	t.Helper()
	base.SetupSuite()
	owner := fixtures.LoginOwner(t)
	buyer := fixtures.LoginBuyer(t, owner)
	return buyer.Token, fixtures.CreateAdvertisement(t, owner.Token)
}

// newAdvertisementID создаёт корректный UUID для сценариев с несуществующим объявлением.
func newAdvertisementID() string {
	var id [16]byte
	rand.Read(id[:])
	id[6] = (id[6] & 0x0f) | 0x40
	id[8] = (id[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", id[:4], id[4:6], id[6:8], id[8:10], id[10:])
}

// cleanupFavorite регистрирует удаление записи текущего теста из избранного.
func cleanupFavorite(t *testing.T, token, advertisementID string) {
	t.Helper()
	t.Cleanup(func() {
		t.Logf("Очистка: удаляем объявление %s из избранного", advertisementID)
		favorites.DeleteFavorite(t, token, advertisementID).RequireCleanupStatus(t)
	})
}

// addFavoriteForSetup добавляет объявление в избранное и регистрирует очистку.
func addFavoriteForSetup(t *testing.T, token, advertisementID string) {
	t.Helper()
	cleanupFavorite(t, token, advertisementID)
	t.Logf("Шаг: добавляем объявление %s в избранное", advertisementID)
	favorites.AddFavorite(t, token, advertisementID).RequireStatus(t, http.StatusCreated)
}

// listFavorites получает избранное и проверяет, что ответ содержит весь список.
func listFavorites(t *testing.T, token string) favorites.FavoritesList {
	t.Helper()
	result := favorites.GetFavorites(t, token)
	result.RequireStatus(t, http.StatusOK)
	list := assertions.Decode[favorites.FavoritesList](t, result)
	require.NotNil(t, list.Total, "список избранного не содержит поле total")
	require.Len(t, list.Items, *list.Total, "ответ должен содержать все записи избранного")
	return list
}

// requireFavoriteState проверяет наличие или отсутствие объявления в избранном.
func requireFavoriteState(t *testing.T, token string, ad *fixtures.PublishedAdvertisement, expectedPresent bool) {
	t.Helper()
	t.Logf("Шаг: проверяем наличие %s в GET /favorites, ожидаем %t", ad.ID, expectedPresent)
	list := listFavorites(t, token)
	found := slices.ContainsFunc(list.Items, func(item favorites.Favorite) bool {
		return item.Advertisement.ID == ad.ID
	})
	require.Equal(t, expectedPresent, found, "состояние избранного для объявления %s (%q)", ad.ID, ad.Title)
	t.Logf("Наличие объявления в избранном подтверждено: %t", found)
}
