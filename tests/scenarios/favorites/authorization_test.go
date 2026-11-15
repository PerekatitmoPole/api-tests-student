package favorites

import (
	"testing"

	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
	"api-tests-template/tests/internal/manager"
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
				result := manager.GetFavorites(t, tc.Token)

				// Assert
				result.RequireUnauthorized(t, tc.Token)
			})
			t.Run("add", func(t *testing.T) {
				// Act
				t.Log("Шаг: отправляем POST /favorites")
				result := manager.AddFavorite(t, tc.Token, id)

				// Assert
				result.RequireUnauthorized(t, tc.Token)
			})
			t.Run("delete", func(t *testing.T) {
				// Act
				t.Log("Шаг: отправляем DELETE /favorites/{ad_id}")
				result := manager.DeleteFavorite(t, tc.Token, id)

				// Assert
				result.RequireUnauthorized(t, tc.Token)
			})
			t.Run("bulk", func(t *testing.T) {
				// Act
				t.Log("Шаг: отправляем POST /favorites/bulk")
				result := manager.AddFavoritesBulk(t, tc.Token, []string{id})

				// Assert
				result.RequireUnauthorized(t, tc.Token)
			})
		})
	}
}
