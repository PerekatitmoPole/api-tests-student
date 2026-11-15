package advertisementCreate

import (
	"testing"

	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
	"api-tests-template/tests/internal/manager"
)

func TestCreateAdvertisementRequiresAuthorization(t *testing.T) {
	// Arrange
	base.SetupSuite()

	for _, tc := range fixtures.UnauthorizedCases() {
		t.Run(tc.Name, func(t *testing.T) {
			request := fixtures.NewAdvertisementRequest(t)

			// Act
			t.Logf("Шаг: отправляем валидное объявление, вариант авторизации %s", tc.Name)
			result := manager.CreateAdvertisement(t, tc.Token, request)

			// Assert
			result.RequireUnauthorized(t, tc.Token)
		})
	}
}
