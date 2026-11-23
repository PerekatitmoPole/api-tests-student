package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

func TestFavoritesRejectUnknownAdvertisement(t *testing.T) {
	// Arrange
	base.SetupSuite()
	user := fixtures.LoginOwner(t)

	cases := []struct {
		name    string
		request func(*testing.T, string, string) assertions.Result
	}{
		{name: "add", request: favorites.AddFavorite},
		{name: "delete", request: favorites.DeleteFavorite},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			id := newAdvertisementID()

			// Act
			t.Logf("Шаг: выполняем операцию %s для несуществующего объявления %s", tc.name, id)
			result := tc.request(t, user.Token, id)

			// Assert
			result.RequireError(t, http.StatusNotFound, "(?i)not found|does not exist")
		})
	}
}
