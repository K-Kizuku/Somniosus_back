package ulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateID() (string, error) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
