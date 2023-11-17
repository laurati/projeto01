package utils

import (
	"context"
	"math/rand"

	"github.com/oklog/ulid/v2"
)

func GenerateUlid(ctx context.Context) (ulid.ULID, error) {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(rand.Int63())), 0)
	ulidId, err := ulid.New(ulid.Now(), entropy)
	if err != nil {
		return ulidId, err
	}
	return ulidId, nil
}
