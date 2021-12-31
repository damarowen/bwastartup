package campaign

import (
	"errors"
	"gorm.io/gorm"
)

type ICampaignRepository interface {
	FindByUserId(id int) ([]Campaign, error)
	FindAll() ([]Campaign, error)
}

type CampaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) ICampaignRepository {
	return &CampaignRepository{db}
}
func (r *CampaignRepository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return  campaigns, nil
}


func (r *CampaignRepository) FindByUserId(id int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", id).Preload("CampaignImages", "campaign_images.is_primary = 1").Take(&campaigns).Error
	if err != nil {
		return campaigns, errors.New("Campaign not found")
	}
	return campaigns, nil
}
