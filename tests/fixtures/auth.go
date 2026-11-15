package fixtures

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
)

func LoginOwner(t *testing.T) models.LoginOkResponse {
	t.Helper()
	return login(t, "TEST_LOGIN", "TEST_PASSWORD")
}

func LoginBuyer(t *testing.T, owner models.LoginOkResponse) models.LoginOkResponse {
	t.Helper()
	buyer := login(t, "TEST_BUYER_LOGIN", "TEST_BUYER_PASSWORD")
	require.NotEqual(t, owner.User.Id, buyer.User.Id, "в .env.override нужны два разных аккаунта")
	return buyer
}

func login(t *testing.T, loginEnv, passwordEnv string) models.LoginOkResponse {
	t.Helper()
	t.Logf("Подготовка: авторизация аккаунта из %s", loginEnv)
	username := os.Getenv(loginEnv)
	password := os.Getenv(passwordEnv)
	require.NotEmpty(t, username, "задайте %s в .env.override", loginEnv)
	require.NotEmpty(t, password, "задайте %s в .env.override", passwordEnv)

	user := auth.Login(t, username, password)
	require.NotEmpty(t, user.Token, "авторизация не вернула токен")
	require.NotEmpty(t, user.User.Id, "авторизация не вернула ID пользователя")
	return user
}

type AuthorizationCase struct {
	Name  string
	Token string
}

func UnauthorizedCases() []AuthorizationCase {
	return []AuthorizationCase{
		{Name: "missing_token"},
		{Name: "invalid_token", Token: "incorrect_token"},
	}
}
