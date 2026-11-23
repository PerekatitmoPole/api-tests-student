package advertisementCreate

import (
	"net/http"
	"slices"
	"strings"
	"testing"

	"api-tests-template/internal/managers/advertisements/models"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

const (
	tooLongTitleLength = 256
	tooManyPhotos      = 4
	photoField         = "photo"
)

func TestCreateAdvertisementRejectsInvalidFields(t *testing.T) {
	// Arrange
	base.SetupSuite()
	owner := fixtures.LoginOwner(t)

	cases := []struct {
		name       string
		field      string
		value      string
		omit       bool
		photoCount int
		message    string
	}{
		{name: "missing_title", field: models.TitleField, omit: true},
		{name: "empty_title", field: models.TitleField, value: ""},
		{name: "missing_description", field: models.DescriptionField, omit: true},
		{name: "missing_photos", field: photoField, photoCount: 0},
		{name: "too_many_photos", field: photoField, photoCount: tooManyPhotos},
		{name: "title_too_long", field: models.TitleField, value: strings.Repeat("a", tooLongTitleLength)},
		{name: "negative_price", field: models.PriceField, value: "-1"},
		{name: "non_numeric_price", field: models.PriceField, value: "abc", message: "^cannot convert string value abc to float$"},
		{name: "negative_quantity", field: models.QuantityField, value: "-1"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			request := fixtures.NewAdvertisementRequest()
			switch {
			case tc.field == photoField:
				request.Photos = slices.Repeat(request.Photos[:1], tc.photoCount)
			case tc.omit:
				delete(request.Fields, tc.field)
			default:
				request.Fields[tc.field] = tc.value
			}
			message := tc.message
			if message == "" {
				message = "(?i)" + tc.field
			}

			// Act
			t.Logf("Шаг: отправляем POST /advertisement, вариант %s", tc.name)
			result, _ := fixtures.PostAdvertisement(t, owner.Token, request)

			// Assert
			result.RequireError(t, http.StatusBadRequest, message)
		})
	}
}
