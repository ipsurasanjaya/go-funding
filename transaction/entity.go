package transaction

import (
	"crowdfunding/campaign"
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
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
