package account

type Authentication struct {
	id             string
	hashedPassword string
}

func NewAuthentication(id, hashedPassword string) *Authentication {
	return &Authentication{
		id:             id,
		hashedPassword: hashedPassword,
	}
}

func (a *Authentication) ID() string {
	return a.id
}

func (a *Authentication) HashedPassword() string {
	return a.hashedPassword
}
