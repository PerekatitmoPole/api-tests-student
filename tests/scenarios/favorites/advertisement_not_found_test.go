package favorites

import (
	"net/http"
	"testing"

	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
	"api-tests-template/tests/internal/manager"
)

func TestFavoritesRejectUnknownAdvertisement(t *testing.T) {
	// Arrange
	base.SetupSuite()
	user := fixtures.LoginOwner(t)

	cases := []struct {
		name    string
		request func(*testing.T, string, string) manager.Result
	}{
		{name: "add", request: manager.AddFavorite},
		{name: "delete", request: manager.DeleteFavorite},
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
