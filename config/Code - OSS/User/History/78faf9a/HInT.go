package tests

import (
	"sso/tests/suite"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

const (
	appID     = 1             // ID приложения, которое мы создали миграцией
	appSecret = "test-secret" // Секретный ключ приложения
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t) // Создаём Suite

	const passDefaultLen = 10

	email := gofakeit.Email()
	pass := randomFakePassword()

	// TODO: Сделать нужные запросы

	// TODO: Проверить результаты
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, passDefaultLen)
}
