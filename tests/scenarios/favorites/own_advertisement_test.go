package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

func TestFavoritesRejectOwnAdvertisement(t *testing.T) {
	// Arrange
	base.SetupSuite()
	owner := fixtures.LoginOwner(t)
	ad := fixtures.CreateAdvertisement(t, owner.Token)
	cleanupFavorite(t, owner.Token, ad.ID)

	// Act
	t.Logf("Шаг: владелец добавляет своё объявление %s в избранное", ad.ID)
	result := favorites.AddFavorite(t, owner.Token, ad.ID)

	// Assert
	result.RequireError(t, http.StatusForbidden, `(?i)\bown\b|owner|belongs`)
}
