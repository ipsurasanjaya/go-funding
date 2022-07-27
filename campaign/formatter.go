package campaign

type FormatCampaign struct {
	CurrentAmount int    `json:"current_amount"`
	FileName      string `json:"file_name"`
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Summary       string `json:"summary"`
	Slug          string `json:"slug"`
	TargetAmmount int    `json:"target_amount"`
	UserID        int    `json:"user_id"`
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

	newCampaign.CurrentAmount = campaign.CurrentAmount
	newCampaign.FileName = ""
	newCampaign.ID = campaign.ID
	newCampaign.Name = campaign.Name
	newCampaign.Summary = campaign.Summary
	newCampaign.Slug = campaign.Slug
	newCampaign.TargetAmmount = campaign.TargetAmount
	newCampaign.UserID = campaign.UserID

	if len(campaign.CampaignImages) != 0 {
		newCampaign.FileName = campaign.CampaignImages[0].FileName
	}

	return newCampaign
}
