package fixtures

import (
	"bytes"
	"image"
	"image/png"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/utils"
	"api-tests-template/tests/internal/manager"
)

type PublishedAdvertisement struct {
	manager.Advertisement
	ownerToken string
	deleted    bool
}

func (ad *PublishedAdvertisement) Delete(t *testing.T) manager.Result {
	t.Helper()
	result := manager.DeleteAdvertisement(t, ad.ownerToken, ad.ID)
	if result.StatusCode == http.StatusNoContent {
		ad.deleted = true
	}
	return result
}

func NewAdvertisementRequest(t *testing.T) manager.CreateAdvertisementRequest {
	t.Helper()
	var photo bytes.Buffer
	require.NoError(t, png.Encode(&photo, image.NewGray(image.Rect(0, 0, 32, 24))),
		"не удалось сформировать тестовое PNG")
	return manager.CreateAdvertisementRequest{
		Fields: map[string]string{
			manager.TitleField:       "QA API " + utils.RandomString(20),
			manager.DescriptionField: "Тестовое объявление для проверки TestBoard API",
			manager.PriceField:       "249.50",
			manager.QuantityField:    "3",
		},
		Photos: [][]byte{photo.Bytes()},
	}
}

func PostAdvertisement(t *testing.T, token string, request manager.CreateAdvertisementRequest) (manager.Result, *PublishedAdvertisement) {
	t.Helper()
	result := manager.CreateAdvertisement(t, token, request)
	if result.StatusCode != http.StatusCreated {
		return result, nil
	}

	ad := &PublishedAdvertisement{Advertisement: manager.Decode[manager.Advertisement](t, result), ownerToken: token}
	require.NotEmpty(t, ad.ID, "ответ о создании не содержит ID для удаления объявления")
	t.Cleanup(func() {
		if ad.deleted {
			return
		}
		t.Logf("Очистка: удаляем объявление %s", ad.ID)
		manager.DeleteAdvertisement(t, token, ad.ID).RequireCleanupStatus(t)
	})
	return result, ad
}

func CreateAdvertisement(t *testing.T, token string) *PublishedAdvertisement {
	t.Helper()
	request := NewAdvertisementRequest(t)
	t.Logf("Подготовка: публикуем объявление %q", request.Fields[manager.TitleField])
	result, ad := PostAdvertisement(t, token, request)
	result.RequireStatus(t, http.StatusCreated)
	require.Equal(t, request.Fields[manager.TitleField], ad.Title, "название в ответе о создании")
	return ad
}
