package campaign

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/mashingan/smapping"
	"log"
)

type ICampaignService interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(dto DtoCampaignDetailById) (Campaign, error)
	CreateCampaign(dto DtoCreateCampaign) (Campaign, error)
	UpdateCampaign(id DtoCampaignDetailById, dto DtoUpdateCampaign) (Campaign, error)
	SaveCampaignImage(input DtoCreateCampaignImage, fileLocation string) (CampaignImage, error)
}

type CampaignService struct {
	CampaignRepository ICampaignRepository
}

func NewCampaignService(CampaignRepo ICampaignRepository) ICampaignService {
	return &CampaignService{CampaignRepo}
}

func (s *CampaignService) GetCampaigns(userId int) ([]Campaign, error) {

	if userId != 0 {
		c, err := s.CampaignRepository.FindByUserId(userId)
		if err != nil {
			return c, err
		}
		return c, nil

	}

	c, err := s.CampaignRepository.FindAll()
	if err != nil {
		return c, err
	}

	return c, nil
}

func (s *CampaignService) GetCampaignById(dto DtoCampaignDetailById) (Campaign, error) {
	_ID, _ := uuid.Parse(dto.ID)
	c, err := s.CampaignRepository.FindById(_ID)

	if err != nil {
		return c, err
	}

	return c, nil
}

func (s *CampaignService) CreateCampaign(input DtoCreateCampaign) (Campaign, error) {

	campaign := Campaign{}

	err := smapping.FillStruct(&campaign, smapping.MapFields(&input))

	campaign.UserId = input.User.ID
	slugify := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugify)

	c, err := s.CampaignRepository.SaveCampaign(campaign)

	if err != nil {
		return c, err
	}

	return c, nil
}

func (s *CampaignService) UpdateCampaign(id DtoCampaignDetailById, input DtoUpdateCampaign) (Campaign, error) {
	_ID, _ := uuid.Parse(id.ID)
	c, err := s.CampaignRepository.FindById(_ID)

	if err != nil {
		return c, err
	}
	if c.UserId != input.User.ID {
		return c, errors.New("not authorize, different user")
	}

	c.Name = input.Name
	c.Description = input.Description
	c.ShortDescription = input.ShortDescription
	c.GoalAmount = input.GoalAmount
	c.Perks = input.Perks
	slugify := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	c.Slug = slug.Make(slugify)

	data, errs := s.CampaignRepository.UpdateCampaign(c)

	if errs != nil {
		return data, err
	}

	return data, nil

}

func (s *CampaignService) SaveCampaignImage(input DtoCreateCampaignImage, fileLocation string) (CampaignImage, error) {

	_ID, _ := uuid.Parse(input.CampaignId)
	c, err := s.CampaignRepository.FindById(_ID)
	if err != nil {
		return CampaignImage{}, err
	}

	if c.UserId != input.User.ID {
		return CampaignImage{}, errors.New("not authorize, different user")
	}

	isPrimary := 0
	//if is true
	if input.IsPrimary {
		isPrimary = 1

		//reset image is primary
		ok, err := s.CampaignRepository.SetImagesNonPrimary(_ID)
		if !ok || err != nil{
			log.Panicln("Set images error")
		}
	}

	campaignImage := CampaignImage{
		CampaignId : _ID,
		IsPrimary : isPrimary,
		FileName : fileLocation,
	}

	newCampaignImage, err := s.CampaignRepository.UploadImagesCampaign(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}
