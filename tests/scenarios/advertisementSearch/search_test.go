package advertisementSearch

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/assertions"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/utils"
	base "api-tests-template/tests"
)

const (
	existingAdTitle = "QAAvitoTestForPracticeLab"
	searchBeginning = "QAAvitoTest"
	searchMiddle    = "TestForPractice"
	searchEnd       = "PracticeLab"
)

type TestSuite struct {
	suite.Suite
	token string
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	t := s.T()
	base.Precondition("Проверяем данные существующего объявления")
	require.NotEmpty(t, existingAdTitle, "укажите полное название объявления в existingAdTitle")
	require.NotEmpty(t, searchBeginning, "укажите начало названия в searchBeginning")
	require.NotEmpty(t, searchMiddle, "укажите подстроку из середины названия в searchMiddle")
	require.NotEmpty(t, searchEnd, "укажите конец названия в searchEnd")
	s.token = auth.Login(t, os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD")).Token
	require.NotEmpty(t, s.token, "авторизация не вернула токен")
}

func (s *TestSuite) TestSearchByTitle() {
	t := s.T()
	list := advertisements.GetAdvertisementsList(t, s.token, searchMiddle)
	require.NotEmpty(t, list.Items, "поиск по %q вернул пустой список", searchMiddle)
	assertions.ContainsAdvertisementByTitle(t, list.Items, existingAdTitle)
}

func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
		name  string
		query string
	}{
		{name: "начало названия", query: searchBeginning},
		{name: "середина названия", query: searchMiddle},
		{name: "конец названия", query: searchEnd},
		{name: "нижний регистр", query: strings.ToLower(searchMiddle)},
		{name: "верхний регистр", query: strings.ToUpper(searchMiddle)},
		{name: "смешанный регистр", query: utils.MixCase(searchMiddle)},
	}
	for _, tc := range cases {
		s.Run(tc.name, func() {
			t := s.T()
			list := advertisements.GetAdvertisementsList(t, s.token, tc.query)
			assertions.ContainsAdvertisementByTitle(t, list.Items, existingAdTitle)
		})
	}
}
