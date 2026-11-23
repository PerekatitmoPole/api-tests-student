package favorites

import (
	"testing"

	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

func TestFavoritesRequireAuthorization(t *testing.T) {
	// Arrange
	base.SetupSuite()

	for _, tc := range fixtures.UnauthorizedCases() {
		t.Run(tc.Name, func(t *testing.T) {
			id := newAdvertisementID()
			t.Run("list", func(t *testing.T) {
				// Act
				t.Log("Шаг: отправляем GET /favorites")
				result := favorites.GetFavorites(t, tc.Token)

				// Assert
				result.RequireUnauthorized(t, tc.Token)
			})
			t.Run("add", func(t *testing.T) {
				// Act
				t.Log("Шаг: отправляем POST /favorites")
				result := favorites.AddFavorite(t, tc.Token, id)

				// Assert
				result.RequireUnauthorized(t, tc.Token)
			})
			t.Run("delete", func(t *testing.T) {
				// Act
				t.Log("Шаг: отправляем DELETE /favorites/{ad_id}")
				result := favorites.DeleteFavorite(t, tc.Token, id)

				// Assert
				result.RequireUnauthorized(t, tc.Token)
			})
			t.Run("bulk", func(t *testing.T) {
				// Act
				t.Log("Шаг: отправляем POST /favorites/bulk")
				result := favorites.AddFavoritesBulk(t, tc.Token, []string{id})

				// Assert
				result.RequireUnauthorized(t, tc.Token)
			})
		})
	}
}
