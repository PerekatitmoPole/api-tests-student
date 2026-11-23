package favorites

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

const favoritesLimit = 50

func TestFavoritesRejectAddingBeyondLimit(t *testing.T) {
	// Arrange
	base.SetupSuite()
	owner := fixtures.LoginOwner(t)
	buyer := fixtures.LoginBuyer(t, owner)
	require.Zero(t, *listFavorites(t, buyer.Token).Total,
		"для проверки лимита нужен выделенный аккаунт покупателя с пустым избранным")

	t.Logf("Подготовка: создаём объявления (%d шт.) и добавляем первые %d в избранное", favoritesLimit+1, favoritesLimit)
	adIDs := make([]string, favoritesLimit)
	for i := range adIDs {
		ad := fixtures.CreateAdvertisement(t, owner.Token)
		adIDs[i] = ad.ID
		cleanupFavorite(t, buyer.Token, ad.ID)
	}
	extraAd := fixtures.CreateAdvertisement(t, owner.Token)
	cleanupFavorite(t, buyer.Token, extraAd.ID)

	favorites.AddFavoritesBulk(t, buyer.Token, adIDs).RequireStatus(t, http.StatusOK)
	require.Equal(t, favoritesLimit, *listFavorites(t, buyer.Token).Total, "лимит избранного не достигнут")

	// Act
	t.Logf("Шаг: пытаемся добавить объявление %s сверх лимита", extraAd.ID)
	result := favorites.AddFavorite(t, buyer.Token, extraAd.ID)

	// Assert
	result.RequireError(t, http.StatusBadRequest, "(?i)limit|maximum")
	requireFavoriteState(t, buyer.Token, extraAd, false)
}
