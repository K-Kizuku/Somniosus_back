package ulid

import (
	"log"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateID() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	if err != nil {
		log.Fatalln("can't generate ulid")
		return ""
	}
	return id.String()
}
