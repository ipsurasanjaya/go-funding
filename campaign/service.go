package campaign

import (
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(ID int) (Campaign, error)
	CreateCampaign(campaignInput CampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaignByID(ID int) (Campaign, error) {
	campaign, err := s.repository.FindByID(ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(campaignInput CampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = campaignInput.Name
	campaign.Summary = campaignInput.Summary
	campaign.Description = campaignInput.Description
	campaign.TargetAmount = campaignInput.TargetAmount
	campaign.Perks = campaignInput.Perks

	slugString := fmt.Sprintf("%s %d", campaignInput.Name, campaignInput.User.Id)
	campaign.Slug = slug.Make(slugString)
	campaign.UserID = campaignInput.User.Id

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
