package transaction

import "gorm.io/gorm"

type Repository interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *repository {
	return &repository{DB: DB}
}

func (r *repository) GetByCampaignID(campaignID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.DB.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
