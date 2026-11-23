package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/auth"
	"api-tests-template/internal/managers/auth/models"
)

// Login выполняет авторизацию пользователя и возвращает токен с данными пользователя, ожидает 200
func Login(t *testing.T, username, password string) models.LoginOkResponse {
	t.Helper()
	require.NotEmpty(t, username, "Отсутствует логин пользователя")
	require.NotEmpty(t, password, "Отсутствует пароль пользователя")

	request := models.LoginRequest{
		Email:    username,
		Password: password,
	}

	loginRawResponse := auth.HttpPostAuthLogin(t, request)
	require.NotNil(t, loginRawResponse, "POST /auth/login не вернул HTTP-ответ")
	defer loginRawResponse.Body.Close()
	require.Equal(t, http.StatusOK, loginRawResponse.StatusCode, "HTTP-статус POST /auth/login")

	respArray, errResponse := io.ReadAll(loginRawResponse.Body)
	require.NoError(t, errResponse, "не удалось прочитать ответ POST /auth/login")

	loginResponse := models.LoginOkResponse{}
	err := json.Unmarshal(respArray, &loginResponse)
	require.NoError(t, err, "не удалось разобрать JSON ответа POST /auth/login")

	return loginResponse
}
