package advertisementSearch

import (
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/assertions"
	catalog "api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
	"api-tests-template/tests/fixtures"
)

const (
	existingAdID          = "0b1bfd39-3066-4312-8842-52fe3e8306f0"
	existingAdDescription = "QAAvitoTestForPracticeLab"
	existingPhotoID       = "4941db61-6bd4-493b-9365-913a6e603865"
)

func TestExistingAdvertisement(t *testing.T) {
	base.SetupSuite()
	user := fixtures.LoginOwner(t)

	t.Run("search_by_full_title", func(t *testing.T) {
		t.Logf("Шаг: ищем объявление по полному названию %q", existingAdTitle)
		list := catalog.GetAdvertisementsList(t, user.Token, existingAdTitle)
		assertions.ContainsAdvertisementByTitle(t, list.Items, existingAdTitle)
		t.Logf("Объявление %q найдено в каталоге", existingAdTitle)
	})

	t.Run("card", func(t *testing.T) {
		t.Logf("Шаг: получаем карточку существующего объявления %s", existingAdID)
		ad := catalog.GetAdvertisement(t, user.Token, existingAdID)
		require.Equal(t, existingAdID, ad.ID, "ID объявления")
		require.Equal(t, existingAdTitle, ad.Title, "название объявления")
		require.Equal(t, existingAdDescription, ad.Description, "описание объявления")
		require.Equal(t, "null", string(ad.Price), "цена в тестовом объявлении не задана")
		require.Equal(t, "null", string(ad.Quantity), "количество в тестовом объявлении не задано")
		require.Len(t, ad.Photos, 1, "количество фотографий в карточке")
		require.Equal(t, existingPhotoID, ad.Photos[0].ID, "ID фото в карточке")
		t.Log("Поля карточки совпадают с ожидаемыми значениями")
	})

	t.Run("photos", func(t *testing.T) {
		t.Logf("Шаг: получаем фото существующего объявления %s", existingAdID)
		photos := catalog.GetAdvertisementPhotos(t, user.Token, existingAdID)
		require.Len(t, photos, 1, "количество фотографий в ответе API")
		require.Equal(t, existingPhotoID, photos[0].ID, "ID фотографии")
		require.NotEmpty(t, photos[0].URL, "API не вернул ссылку на фото")
		t.Logf("API вернул ожидаемое фото %s и ссылку", existingPhotoID)
	})
}
