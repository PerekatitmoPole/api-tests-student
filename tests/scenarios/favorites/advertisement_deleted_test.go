package favorites

import (
	"net/http"
	"testing"

	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

func TestDeletingAdvertisementRemovesItFromFavorites(t *testing.T) {
	// Arrange
	base.SetupSuite()
	owner := fixtures.LoginOwner(t)
	buyer := fixtures.LoginBuyer(t, owner)
	ad := fixtures.CreateAdvertisement(t, owner.Token)

	addFavoriteForSetup(t, buyer.Token, ad.ID)
	requireFavoriteState(t, buyer.Token, ad, true)

	// Act
	t.Logf("Шаг: владелец удаляет объявление %s", ad.ID)
	result := ad.Delete(t)

	// Assert
	result.RequireStatus(t, http.StatusNoContent)
	requireFavoriteState(t, buyer.Token, ad, false)
}
