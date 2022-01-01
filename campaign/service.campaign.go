package campaign

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"errors"
)

type ICampaignService interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(dto DtoCampaignDetailById) (Campaign, error)
	CreateCampaign(dto DtoCreateCampaign) (Campaign, error)
	UpdateCampaign(id DtoCampaignDetailById, dto DtoUpdateCampaign) (Campaign, error)
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

	data := Campaign{
		Name:             input.Name,
		Description:      input.Description,
		ShortDescription: input.ShortDescription,
		GoalAmount:       input.GoalAmount,
		Perks:            input.Perks,
	}
	data.UserId = input.User.ID
	slugify := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	data.Slug = slug.Make(slugify)

	c, err := s.CampaignRepository.SaveCampaign(data)

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
	fmt.Println(c.UserId != input.User.ID, ",,,<<<")
	if c.UserId != input.User.ID{
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
