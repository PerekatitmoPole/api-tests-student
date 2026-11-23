package favorites

import (
	"net/http"
	"testing"
)

func TestDeletingAdvertisementRemovesItFromFavorites(t *testing.T) {
	// Arrange
	buyerToken, ad := setupFavoriteAdvertisement(t)

	addFavoriteForSetup(t, buyerToken, ad.ID)
	requireFavoriteState(t, buyerToken, ad, true)

	// Act
	t.Logf("Шаг: владелец удаляет объявление %s", ad.ID)
	result := ad.Delete(t)

	// Assert
	result.RequireStatus(t, http.StatusNoContent)
	requireFavoriteState(t, buyerToken, ad, false)
}
