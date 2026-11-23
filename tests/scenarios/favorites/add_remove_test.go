package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/favorites"
)

func TestAddAndRemoveFavorite(t *testing.T) {
	// Arrange
	buyerToken, ad := setupFavoriteAdvertisement(t)
	cleanupFavorite(t, buyerToken, ad.ID)

	// Act
	t.Logf("Шаг: добавляем объявление %s в избранное", ad.ID)
	added := favorites.AddFavorite(t, buyerToken, ad.ID)

	// Assert
	added.RequireStatus(t, http.StatusCreated)
	requireFavoriteState(t, buyerToken, ad, true)

	// Act
	t.Logf("Шаг: удаляем объявление %s из избранного покупателя", ad.ID)
	removed := favorites.DeleteFavorite(t, buyerToken, ad.ID)

	// Assert
	removed.RequireStatus(t, http.StatusNoContent)
	requireFavoriteState(t, buyerToken, ad, false)
}
