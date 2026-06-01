package tests

import "github.com/brianvoe/gofakeit/v6"

func randomUserID() int64 {
	return int64(gofakeit.Number(10000, 99999))
}
