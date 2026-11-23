package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/favorites"
)

func TestFavoritesRejectDuplicateAdvertisement(t *testing.T) {
	// Arrange
	buyerToken, ad := setupFavoriteAdvertisement(t)
	addFavoriteForSetup(t, buyerToken, ad.ID)

	// Act
	t.Logf("Шаг: повторно добавляем объявление %s", ad.ID)
	result := favorites.AddFavorite(t, buyerToken, ad.ID)

	// Assert
	result.RequireError(t, http.StatusConflict, "(?i)already|duplicate")
}
