package fixtures

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
)

// LoginOwner авторизует владельца объявления с учётными данными из окружения.
func LoginOwner(t *testing.T) models.LoginOkResponse {
	t.Helper()
	return login(t, "владелец", "TEST_LOGIN", "TEST_PASSWORD")
}

// LoginBuyer авторизует покупателя и проверяет, что его аккаунт отличается от владельца.
func LoginBuyer(t *testing.T, owner models.LoginOkResponse) models.LoginOkResponse {
	t.Helper()
	buyer := login(t, "покупатель", "TEST_BUYER_LOGIN", "TEST_BUYER_PASSWORD")
	require.NotEqual(t, owner.User.Id, buyer.User.Id, "TEST_LOGIN и TEST_BUYER_LOGIN должны указывать на разные аккаунты")
	return buyer
}

// login читает параметры указанной роли из окружения и проверяет результат входа.
func login(t *testing.T, role, loginEnv, passwordEnv string) models.LoginOkResponse {
	t.Helper()
	t.Logf("Подготовка: вход в TestBoard для роли %s (%s)", role, loginEnv)
	username := os.Getenv(loginEnv)
	password := os.Getenv(passwordEnv)
	require.NotEmpty(t, username, "роль %s: задайте %s в .env.override", role, loginEnv)
	require.NotEmpty(t, password, "роль %s: задайте %s в .env.override", role, passwordEnv)

	user := auth.Login(t, username, password)
	require.NotEmpty(t, user.Token, "роль %s: авторизация не вернула токен", role)
	require.NotEmpty(t, user.User.Id, "роль %s: авторизация не вернула ID пользователя", role)
	return user
}

type AuthorizationCase struct {
	Name  string
	Token string
}

// UnauthorizedCases возвращает параметры запросов без действующей авторизации.
func UnauthorizedCases() []AuthorizationCase {
	return []AuthorizationCase{
		{Name: "missing_token"},
		{Name: "invalid_token", Token: "incorrect_token"},
	}
}
