package campaign

type DtoCampaign struct {
	ID             int `json:"id"`
	UserId           int `json:"user_id"`
	Name     string `json:"name"`
	ShortDescription          string `json:"short_description"`
	Slug  string`json:"slug"`
	ImageUrl   string `json:"image_url"`
	GoalAmount           int `json:"goal_amount"`
	CurrentAmount           int `json:"current_amount"`
}

func MappingResponseCampaign(campaigns []Campaign) []DtoCampaign {

	var storeCampaign[]DtoCampaign

	for _, c := range campaigns {

		var _imageUrl string
		if (len(c.CampaignImages)) > 0 {
			_imageUrl = 	c.CampaignImages[0].FileName
		}

		data := DtoCampaign{
			ID:         c.ID,
			UserId: c.UserId,
			Name: c.Name,
			ShortDescription: c.ShortDescription,
			Slug: c.Slug,
			GoalAmount: c.GoalAmount,
			CurrentAmount: c.CurrentAmount,
			ImageUrl: _imageUrl,
		}

		storeCampaign = append(storeCampaign, data)
	}

	return storeCampaign
}


