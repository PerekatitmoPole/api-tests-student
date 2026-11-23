package advertisements

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/require"

	httpclient "api-tests-template/internal/client/http"
)

const (
	advertisementPath       = "/advertisement"
	advertisementPhotosPath = "/advertisements/%s/photos"
)

// HttpPostAdvertisement публикует объявление с полями и фотографиями multipart.
func HttpPostAdvertisement(t *testing.T, token string, fields map[string]string, photos [][]byte) *http.Response {
	t.Helper()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	for name, value := range fields {
		require.NoError(t, writer.WriteField(name, value), "не удалось записать поле %q", name)
	}
	header := textproto.MIMEHeader{
		"Content-Disposition": {`form-data; name="photos"; filename="test-photo.png"`},
		"Content-Type":        {"image/png"},
	}
	for _, photo := range photos {
		part, err := writer.CreatePart(header)
		require.NoError(t, err, "не удалось добавить фото в запрос")
		_, err = part.Write(photo)
		require.NoError(t, err, "не удалось записать содержимое фото")
	}
	require.NoError(t, writer.Close(), "не удалось завершить multipart-запрос")

	return httpclient.NewRequest(token).Post(advertisementPath).
		ContentType(writer.FormDataContentType()).Body(body.String()).Expect(t).End().Response
}

// HttpGetAdvertisement запрашивает карточку по ID объявления.
func HttpGetAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return httpclient.NewRequest(token).Get(advertisementPath).
		Query("id", id).Expect(t).End().Response
}

// HttpDeleteAdvertisement удаляет объявление по его ID.
func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return httpclient.NewRequest(token).Delete(advertisementPath).
		Query("id", id).Expect(t).End().Response
}

// HttpGetAdvertisementPhotos запрашивает фотографии объявления.
func HttpGetAdvertisementPhotos(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return httpclient.NewRequest(token).
		Getf(advertisementPhotosPath, id).Expect(t).End().Response
}
