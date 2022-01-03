package campaign

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ICampaignRepository interface {
	FindByUserId(id int) ([]Campaign, error)
	FindById(id uuid.UUID) (Campaign, error)
	FindAll() ([]Campaign, error)
	SaveCampaign(campaign Campaign) (Campaign, error)
	UpdateCampaign(campaign Campaign) (Campaign, error)
	UploadImagesCampaign(cImage CampaignImage) (CampaignImage, error)
	SetImagesNonPrimary(campaignId uuid.UUID) (bool, error)
}

type CampaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) ICampaignRepository {
	return &CampaignRepository{db}
}

func (r *CampaignRepository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Order("created_at desc").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *CampaignRepository) FindByUserId(id int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Where("user_id = ?", id).Take(&campaigns).Error
	if err != nil {
		return campaigns, errors.New("Campaign not found")
	}
	return campaigns, nil
}

func (r *CampaignRepository) FindById(id uuid.UUID) (Campaign, error) {
	var c Campaign
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", id).Take(&c).Error
	if err != nil {
		return c, errors.New("Campaign not found")
	}
	return c, nil
}

func (r *CampaignRepository) SaveCampaign(c Campaign) (Campaign, error) {
	err := r.db.Create(&c).Error

	if err != nil {
		return c, err
	}
	return c, nil
}

func (r *CampaignRepository) UpdateCampaign(c Campaign) (Campaign, error) {
	err := r.db.Save(&c).Error
	if err != nil {
		return c, err
	}
	return c, nil
}

func (r *CampaignRepository) UploadImagesCampaign(c CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&c).Error

	if err != nil {
		return c, err
	}
	return c, nil
}

func (r *CampaignRepository) SetImagesNonPrimary(campaignId uuid.UUID) (bool, error) {
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignId).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil

}
