package dto

import (
	"time"

	"github.com/K-Kizuku/Somniosus_back/apps/account/app/domain/model/account"
)

type Account struct {
	ID          string                `db:"id"`
	Status      account.AccountStatus `db:"status"`
	DisplayID   string                `db:"display_id"`
	Name        string                `db:"name"`
	Description string                `db:"description"`
	ImageURL    string                `db:"image_url"`
	BirthDay    time.Time             `db:"birth_day"`
	TelNum      string                `db:"tel_num"`
	CreatedAt   *time.Time            `db:"created_at"`
	Updated_at  *time.Time            `db:"updated_at"`
}
