package dto

type Authentication struct {
	ID             string `db:"id"`
	HashedPassword string `db:"hashed_password"`
}
