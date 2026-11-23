package assertions

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type Result struct {
	StatusCode int
	body       []byte
}

type errorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// ReadResponse читает ответ и сохраняет тело для проверок и диагностики.
func ReadResponse(t *testing.T, res *http.Response) Result {
	t.Helper()
	require.NotNil(t, res, "сервер не вернул HTTP-ответ")
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err, "не удалось прочитать тело ответа")
	t.Logf("Ответ API: HTTP %d", res.StatusCode)
	return Result{StatusCode: res.StatusCode, body: body}
}

// RequireStatus сравнивает HTTP-статус и показывает тело при несовпадении.
func (r Result) RequireStatus(t *testing.T, expected int) {
	t.Helper()
	require.Equal(t, expected, r.StatusCode, "тело ответа: %s", r.body)
}

// RequireCleanupStatus принимает успешное удаление или отсутствие записи.
func (r Result) RequireCleanupStatus(t *testing.T) {
	t.Helper()
	require.Contains(t, []int{http.StatusNoContent, http.StatusNotFound}, r.StatusCode,
		"не удалось удалить тестовые данные, тело ответа: %s", r.body)
}

// Decode разбирает JSON из HTTP-ответа в указанную модель.
func Decode[T any](t *testing.T, r Result) T {
	t.Helper()
	var data T
	require.NoError(t, json.Unmarshal(r.body, &data), "не удалось разобрать ответ: %s", r.body)
	return data
}

// RequireError проверяет HTTP-статус и содержание ошибки API.
func (r Result) RequireError(t *testing.T, status int, messagePattern string) {
	t.Helper()
	r.RequireStatus(t, status)
	apiError := Decode[errorResponse](t, r)
	require.NotEmpty(t, apiError.Error, "ответ не содержит код ошибки")
	require.Regexp(t, messagePattern, apiError.Message, "неожиданное сообщение об ошибке")
	t.Logf("Ошибка подтверждена: %s, %s", apiError.Error, apiError.Message)
}

// RequireUnauthorized проверяет ответ для отсутствующего или неверного токена.
func (r Result) RequireUnauthorized(t *testing.T, token string) {
	t.Helper()
	r.RequireStatus(t, http.StatusUnauthorized)
	apiError := Decode[errorResponse](t, r)
	expectedMessage := "Authorization header required"
	if token != "" {
		expectedMessage = "Invalid or expired token"
	}
	require.Equal(t, "unauthorized", apiError.Error, "неожиданный код ошибки авторизации")
	require.Equal(t, expectedMessage, apiError.Message, "неожиданное сообщение об ошибке авторизации")
	t.Logf("Ошибка авторизации подтверждена: %s", apiError.Message)
}
