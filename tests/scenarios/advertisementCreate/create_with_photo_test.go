package advertisementCreate

import (
	"net/http"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

func TestCreateAdvertisementWithAllParameters(t *testing.T) {
	// Arrange
	base.SetupSuite()
	owner := fixtures.LoginOwner(t)
	request := fixtures.NewAdvertisementRequest()
	title := request.Fields[models.TitleField]

	// Act
	t.Logf("Шаг: публикуем объявление %q со всеми полями и фото", title)
	result, created := fixtures.PostAdvertisement(t, owner.Token, request)

	// Assert
	result.RequireStatus(t, http.StatusCreated)
	t.Logf("Объявление создано: %s", created.ID)

	t.Run("card", func(t *testing.T) {
		t.Logf("Шаг: проверяем сохранённые поля объявления %s", created.ID)
		card := advertisements.GetAdvertisement(t, owner.Token, created.ID)
		require.Equal(t, created.ID, card.ID, "ID в карточке объявления")
		require.Equal(t, owner.User.Id, card.UserID, "владелец в карточке объявления")
		require.Equal(t, title, card.Title, "название в карточке объявления")
		require.Equal(t, request.Fields[models.DescriptionField], card.Description, "описание в карточке объявления")
		require.JSONEq(t, request.Fields[models.PriceField], string(card.Price), "цена в карточке объявления")
		require.JSONEq(t, request.Fields[models.QuantityField], string(card.Quantity), "количество в карточке объявления")
		require.Equal(t, created.Photos, card.Photos, "фотографии в карточке объявления")
		t.Log("Карточка вернула все переданные поля")
	})

	t.Run("search_by_full_title", func(t *testing.T) {
		t.Logf("Шаг: ищем объявление по полному названию %q", title)
		items := advertisements.GetAdvertisementsList(t, owner.Token, title).Items
		found := slices.ContainsFunc(items, func(item models.Advertisement) bool {
			return item.ID == created.ID && item.Title == title
		})
		require.True(t, found, "поиск по %q не вернул созданное объявление %s", title, created.ID)
		t.Logf("Поиск вернул объявление %s с точным названием", created.ID)
	})

	t.Run("photos", func(t *testing.T) {
		t.Logf("Шаг: запрашиваем список фото объявления %s", created.ID)
		require.Len(t, created.Photos, len(request.Photos), "число фото в ответе о создании")
		photos := advertisements.GetAdvertisementPhotos(t, owner.Token, created.ID)
		require.Equal(t, created.Photos, photos, "GET /advertisements/{id}/photos должен вернуть загруженное фото")
		require.NotEmpty(t, photos[0].ID, "API не вернул ID загруженного фото")
		require.NotEmpty(t, photos[0].URL, "API не вернул ссылку на загруженное фото")
		t.Logf("API вернул фото %s и ссылку", photos[0].ID)
	})
}
