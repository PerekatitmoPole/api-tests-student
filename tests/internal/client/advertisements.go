package client

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	advertisementPath       = "/advertisement"
	advertisementPhotosPath = "/advertisements/%s/photos"
)

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

	return newRequest(token).Post(advertisementPath).
		ContentType(writer.FormDataContentType()).Body(body.String()).Expect(t).End().Response
}

func HttpGetAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return newRequest(token).Get(advertisementPath).
		Query("id", id).Expect(t).End().Response
}

func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return newRequest(token).Delete(advertisementPath).
		Query("id", id).Expect(t).End().Response
}

func HttpGetAdvertisementPhotos(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return newRequest(token).
		Getf(advertisementPhotosPath, id).Expect(t).End().Response
}
