package campaign

type ICampaignService interface {
	GetCampaigns(userId int) ([]Campaign, error)
}

type CampaignService struct {
	CampaingRepository ICampaignRepository
}

func NewCampaignService(CampaingRepo ICampaignRepository) ICampaignService {
	return &CampaignService{CampaingRepo}
}

func (s *CampaignService) GetCampaigns(userId int) ([]Campaign, error) {

	if userId != 0 {
			campaings, err := s.CampaingRepository.FindByUserId(userId)
			if err != nil {
				return campaings, err
			}
			return campaings , nil

	}

	campaings, err := s.CampaingRepository.FindAll()
	if err != nil {
		return campaings, err
	}

return campaings , nil
}