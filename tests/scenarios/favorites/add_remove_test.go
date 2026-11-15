package favorites

import (
	"net/http"
	"testing"

	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
	"api-tests-template/tests/internal/manager"
)

func TestAddAndRemoveFavorite(t *testing.T) {
	// Arrange
	base.SetupSuite()
	owner := fixtures.LoginOwner(t)
	buyer := fixtures.LoginBuyer(t, owner)
	ad := fixtures.CreateAdvertisement(t, owner.Token)
	cleanupFavorite(t, buyer.Token, ad.ID)

	// Act
	t.Logf("Шаг: добавляем объявление %s в избранное", ad.ID)
	added := manager.AddFavorite(t, buyer.Token, ad.ID)

	// Assert
	added.RequireStatus(t, http.StatusCreated)
	requireFavoriteState(t, buyer.Token, ad, true)

	// Act
	t.Logf("Шаг: удаляем объявление %s из избранного покупателя", ad.ID)
	removed := manager.DeleteFavorite(t, buyer.Token, ad.ID)

	// Assert
	removed.RequireStatus(t, http.StatusNoContent)
	requireFavoriteState(t, buyer.Token, ad, false)
}
