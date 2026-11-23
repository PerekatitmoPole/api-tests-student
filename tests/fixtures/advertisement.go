package fixtures

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	"api-tests-template/internal/utils"
	"api-tests-template/tests/scenarios/advertisementCreate/testdata"
)

type PublishedAdvertisement struct {
	models.AdvertisementDetails
	ownerToken string
	deleted    bool
}

// Delete удаляет объявление и отмечает успешное удаление для очистки.
func (ad *PublishedAdvertisement) Delete(t *testing.T) assertions.Result {
	t.Helper()
	result := advertisements.DeleteAdvertisement(t, ad.ownerToken, ad.ID)
	if result.StatusCode == http.StatusNoContent {
		ad.deleted = true
	}
	return result
}

// NewAdvertisementRequest создаёт уникальные данные объявления с тестовым фото.
func NewAdvertisementRequest() models.CreateAdvertisementRequest {
	return models.CreateAdvertisementRequest{
		Fields: map[string]string{
			models.TitleField:       "QA API " + utils.RandomString(20),
			models.DescriptionField: "Тестовое объявление для проверки TestBoard API",
			models.PriceField:       "249.50",
			models.QuantityField:    "3",
		},
		Photos: [][]byte{[]byte(testdata.AdvertisementPhoto)},
	}
}

// PostAdvertisement отправляет данные и регистрирует очистку созданного объявления.
func PostAdvertisement(t *testing.T, token string, request models.CreateAdvertisementRequest) (assertions.Result, *PublishedAdvertisement) {
	t.Helper()
	result := advertisements.CreateAdvertisement(t, token, request)
	if result.StatusCode != http.StatusCreated {
		return result, nil
	}

	ad := &PublishedAdvertisement{
		AdvertisementDetails: assertions.Decode[models.AdvertisementDetails](t, result),
		ownerToken:           token,
	}
	require.NotEmpty(t, ad.ID, "ответ о создании не содержит ID для удаления объявления")
	t.Cleanup(func() {
		if ad.deleted {
			return
		}
		t.Logf("Очистка: удаляем объявление %s", ad.ID)
		advertisements.DeleteAdvertisement(t, token, ad.ID).RequireCleanupStatus(t)
	})
	return result, ad
}

// CreateAdvertisement публикует объявление для подготовки сценария.
func CreateAdvertisement(t *testing.T, token string) *PublishedAdvertisement {
	t.Helper()
	request := NewAdvertisementRequest()
	t.Logf("Подготовка: публикуем объявление %q", request.Fields[models.TitleField])
	result, ad := PostAdvertisement(t, token, request)
	result.RequireStatus(t, http.StatusCreated)
	require.Equal(t, request.Fields[models.TitleField], ad.Title, "название в ответе о создании")
	return ad
}
