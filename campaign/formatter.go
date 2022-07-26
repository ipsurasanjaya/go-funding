package campaign

type FormatCampaign struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	Name          string `json:"name"`
	Summary       string `json:"summary"`
	TargetAmmount int    `json:"target_amount"`
	CurrentAmount int    `json:"current_amount"`
	FileName      string `json:"file_name"`
}

func CampaignsFormatter(campaigns []Campaign) []FormatCampaign {
	newCampaigns := []FormatCampaign{}
	for _, campaign := range campaigns {
		newCampaigns = append(newCampaigns, CampaignFormatter(campaign))
	}
	return newCampaigns
}

func CampaignFormatter(campaign Campaign) FormatCampaign {
	var newCampaign FormatCampaign
	newCampaign.ID = campaign.ID
	newCampaign.UserID = campaign.UserID
	newCampaign.Name = campaign.Name
	newCampaign.Summary = campaign.Summary
	newCampaign.TargetAmmount = campaign.TargetAmount
	newCampaign.CurrentAmount = campaign.CurrentAmount
	newCampaign.FileName = ""
	if len(campaign.CampaignImages) != 0 {
		newCampaign.FileName = campaign.CampaignImages[0].FileName
	}

	return newCampaign
}
