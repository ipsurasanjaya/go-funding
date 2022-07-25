package campaign

import "time"

type Campaign struct {
	ID             int
	UserID         int
	Name           string
	TargetAmount   int
	CurrentAmount  int
	Summary        string
	Description    string
	Perks          string
	Slug           string
	BackerCount    int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CampaignImages []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
