package account

import "github.com/K-Kizuku/Somniosus_back/lib/ulid"

func NewID() string {
	id := ulid.GenerateID()
	return id
}
