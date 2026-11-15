package favorites

import (
	"net/http"
	"testing"

	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
	"api-tests-template/tests/internal/manager"
)

func TestFavoritesRejectDuplicateAdvertisement(t *testing.T) {
	// Arrange
	base.SetupSuite()
	owner := fixtures.LoginOwner(t)
	buyer := fixtures.LoginBuyer(t, owner)
	ad := fixtures.CreateAdvertisement(t, owner.Token)
	addFavoriteForSetup(t, buyer.Token, ad.ID)

	// Act
	t.Logf("Шаг: повторно добавляем объявление %s", ad.ID)
	result := manager.AddFavorite(t, buyer.Token, ad.ID)

	// Assert
	result.RequireError(t, http.StatusConflict, "(?i)already|duplicate")
}
