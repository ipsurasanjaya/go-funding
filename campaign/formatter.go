package campaign

import "strings"

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

type FormatCampaignDetails struct {
	CurrentAmount int           `json:"current_amount"`
	FileName      string        `json:"file_name"`
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Summary       string        `json:"summary"`
	Slug          string        `json:"slug"`
	BackerCount   int           `json:"backer_count"`
	TargetAmmount int           `json:"target_amount"`
	UserID        int           `json:"user_id"`
	Perks         []string      `json:"perks"`
	User          formatUser    `json:"user"`
	Images        []formatImage `json:"images"`
}

type formatUser struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type formatImage struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func CampaignDetailsFormatter(campaign Campaign) FormatCampaignDetails {
	campaignDetails := FormatCampaignDetails{
		CurrentAmount: campaign.CurrentAmount,
		FileName:      "",
		ID:            campaign.ID,
		Name:          campaign.Name,
		Summary:       campaign.Summary,
		Slug:          campaign.Slug,
		TargetAmmount: campaign.TargetAmount,
		UserID:        campaign.UserID,
		BackerCount:   campaign.BackerCount,
	}

	if len(campaign.CampaignImages) != 0 {
		campaignDetails.FileName = campaign.CampaignImages[0].FileName
	}
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetails.Perks = perks
	campaignDetails.User = campaignUserFormatter(campaign)
	campaignDetails.Images = campaignImagesFormatter(campaign)

	return campaignDetails
}

func campaignUserFormatter(campaign Campaign) formatUser {
	var user formatUser

	user.Name = campaign.User.Name
	user.ImageUrl = campaign.User.AvatarFileName

	return user
}

func campaignImagesFormatter(campaign Campaign) []formatImage {
	var (
		campaignImage  formatImage
		campaignImages []formatImage
	)

	for _, image := range campaign.CampaignImages {
		campaignImage.ImageUrl = image.FileName
		campaignImage.IsPrimary = image.IsPrimary
		campaignImages = append(campaignImages, campaignImage)
	}
	return campaignImages
}
