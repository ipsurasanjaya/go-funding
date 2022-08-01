package transaction

import (
	"crowdfunding/user"
	"time"
)

type Transaction struct {
	ID         int
	UserID     int
	CampaignID int
	Amount     int
	Status     string
	TrxCode    string
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
