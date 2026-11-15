package favorites

import (
	"net/http"
	"testing"

	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
	"api-tests-template/tests/internal/manager"
)

func TestFavoritesRejectOwnAdvertisement(t *testing.T) {
	// Arrange
	base.SetupSuite()
	owner := fixtures.LoginOwner(t)
	ad := fixtures.CreateAdvertisement(t, owner.Token)
	cleanupFavorite(t, owner.Token, ad.ID)

	// Act
	t.Logf("Шаг: владелец добавляет своё объявление %s в избранное", ad.ID)
	result := manager.AddFavorite(t, owner.Token, ad.ID)

	// Assert
	result.RequireError(t, http.StatusForbidden, `(?i)\bown\b|owner|belongs`)
}
