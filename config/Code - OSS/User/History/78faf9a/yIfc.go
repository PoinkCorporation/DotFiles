package tests

import (
	"testing"

	"sso/tests/suite"

	ssov1 "sso/gen/go/sso"

	"github.com/brianvoe/gofakeit/v6"
)

const (
	appID  = 1
	appKey = "test-secret"

	passDefaultLen = 10
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	email := gofakeit.Email()
	pass := randomFakePassword()

	// Сначала зарегистрируем нового пользователя, которого будем логинить
	respReg, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
		Email:    email,
		Password: pass,
	})
	// Это вспомогательный запрос, поэтому делаем лишь минимальные проверки
	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	// А это основная проверка
	respLogin, err := st.AuthClient.Login(ctx, &ssov1.LoginRequest{
		Email:    email,
		Password: pass,
		AppId:    appID,
	})
	require.NoError(t, err)

	// Тщательно проверяем результаты:

	// TODO: Доставём из ответа токен авторизации и проверяем его содержимое
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, passDefaultLen)
}
