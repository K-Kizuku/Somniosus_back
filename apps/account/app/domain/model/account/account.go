package account

import "time"

type AccountStatus int

const (
	Progressing AccountStatus = iota
	General
	Official
	Subscriber
)

func (a AccountStatus) String() string {
	switch a {
	case Progressing:
		return "progressing"
	case General:
		return "general"
	case Official:
		return "official"
	case Subscriber:
		return "subscriber"
	default:
		return "Unknown"
	}
}

type Account struct {
	id          string
	status      AccountStatus
	displayID   string
	name        string
	description string
	imageURL    string
	birthDay    time.Time
	telNum      string
	websiteURL  string
}

func NewAccount(status AccountStatus, displayID, name, description, imageURL, telNum string, birthDay time.Time) *Account {
	return &Account{
		id:          NewID(),
		status:      status,
		displayID:   displayID,
		name:        name,
		description: description,
		imageURL:    imageURL,
		birthDay:    birthDay,
		telNum:      telNum,
	}
}

func Assign(status AccountStatus, id, displayID, name, description, imageURL, telNum string, birthDay time.Time) *Account {
	return &Account{
		id:          id,
		status:      status,
		displayID:   displayID,
		name:        name,
		description: description,
		imageURL:    imageURL,
		birthDay:    birthDay,
		telNum:      telNum,
	}
}

func (a Account) ID() string {
	return a.id
}

func (a Account) Status() AccountStatus {
	return a.status
}

func (a Account) DisplayID() string {
	return a.displayID
}
func (a Account) Name() string {
	return a.name
}

func (a Account) Description() string {
	return a.description
}

func (a Account) ImageURL() string {
	return a.imageURL
}

func (a Account) BirthDay() time.Time {
	return a.birthDay
}
func (a Account) TelNum() string {
	return a.telNum
}

func (a Account) WebsiteURL() string {
	return a.websiteURL
}
