package advertisementCreate

import (
	"testing"

	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

func TestCreateAdvertisementRequiresAuthorization(t *testing.T) {
	// Arrange
	base.SetupSuite()

	for _, tc := range fixtures.UnauthorizedCases() {
		t.Run(tc.Name, func(t *testing.T) {
			request := fixtures.NewAdvertisementRequest()

			// Act
			t.Logf("Шаг: отправляем валидное объявление, вариант авторизации %s", tc.Name)
			result := advertisements.CreateAdvertisement(t, tc.Token, request)

			// Assert
			result.RequireUnauthorized(t, tc.Token)
		})
	}
}
