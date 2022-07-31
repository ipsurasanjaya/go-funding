package campaign

import "crowdfunding/user"

type CampaignInput struct {
	Name         string `json:"name" binding:"required"`
	Summary      string `json:"summary" binding:"required"`
	Description  string `json:"description" binding:"required"`
	TargetAmount int    `json:"target_amount" binding:"required"`
	Perks        string `json:"perks" binding:"required"`
	User         user.User
}

type CampaignImageInput struct {
	IsPrimary  bool `form:"is_primary" binding:"required"`
	CampaignID int  `form:"campaign_id" binding:"required"`
	User       user.User
}
