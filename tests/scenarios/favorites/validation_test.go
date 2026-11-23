package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

func TestFavoritesRejectInvalidAdvertisementID(t *testing.T) {
	// Arrange
	base.SetupSuite()
	user := fixtures.LoginOwner(t)

	cases := []struct {
		name    string
		id      string
		message string
	}{
		{name: "missing_id", message: "(?i)request|advertisement|id"},
		{name: "malformed_id", id: "not-a-uuid", message: "(?i)advertisement|id|uuid"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			t.Logf("Шаг: отправляем POST /favorites, вариант %s", tc.name)
			result := favorites.AddFavorite(t, user.Token, tc.id)

			// Assert
			result.RequireError(t, http.StatusBadRequest, tc.message)
		})
	}
}
