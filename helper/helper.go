package helper

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"fmt"
	"github.com/mashingan/smapping"
	"log"
	"strings"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Errors  interface{} `json:"errors"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj []interface{}

func ApiResponse(status bool, message string, code int, data interface{}, err interface{}) Response {

	//normal
	//var errors []string
	//for _,e := range err.(validator.ValidationErrors){
	//	errors = append(errors, e.Error())
	//}
	//errMessage := gin.H{"ERROR": errors}
	//

	fmt.Printf("type: %T\n", err)

	var dataErr interface{}
	//logic error = " jika tidak ada error
	if len(err.(string)) > 0 {
		dataErr = strings.Split(err.(string), "\n")
	} else {
		dataErr = ""
	}

	Meta := Meta{
		Status:  status,
		Message: message,
		Code:    code,
		Errors:  dataErr,
	}

	jsonResponse := Response{
		Meta: Meta,
		Data: data,
	}

	return jsonResponse
}

func MappingResponseUser(u user.User, token string) user.DtoResponseUser {

	return user.DtoResponseUser{
		ID:         u.ID,
		Name:       u.Name,
		Occupation: u.Occupation,
		Email:      u.Email,
		Token:      token,
	}

}

func MappingResponseCampaign(c campaign.Campaign) campaign.DtoCampaign {

	var _imageUrl string
	if (len(c.CampaignImages)) > 0 {
		_imageUrl = c.CampaignImages[0].FileName
	}
		return campaign.DtoCampaign{
			ID:         c.ID,
			UserId: c.UserId,
			Name: c.Name,
			ShortDescription: c.ShortDescription,
			Slug: c.Slug,
			GoalAmount: c.GoalAmount,
			CurrentAmount: c.CurrentAmount,
			ImageUrl: _imageUrl,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		}

}

func MappingResponseCampaigns(campaigns []campaign.Campaign) []campaign.DtoCampaign {

	var storeCampaign []campaign.DtoCampaign
	var data campaign.DtoCampaign
	for _, c := range campaigns {

		var _imageUrl string
		if (len(c.CampaignImages)) > 0 {
			_imageUrl = c.CampaignImages[0].FileName
		}

		err := smapping.FillStruct(&data, smapping.MapFields(&c))
		if err != nil {
			log.Fatalf("Failed map %v", err)
		}
		data.ImageUrl = _imageUrl

		storeCampaign = append(storeCampaign, data)
	}

	return storeCampaign
}

func MappingResponseDetailCampaign(c campaign.Campaign) campaign.DtoCampaignDetail {

	dto := campaign.DtoCampaignDetail{
		ID:               c.ID,
		Name:             c.Name,
		ShortDescription: c.ShortDescription,
		Description:      c.Description,
		GoalAmount:       c.GoalAmount,
		CurrentAmount:    c.CurrentAmount,
		Slug:             c.Slug,
	}

	var _imageUrl string
	if (len(c.CampaignImages)) > 0 {
		_imageUrl = c.CampaignImages[0].FileName
	}
	dto.ImageUrl = _imageUrl

	var perks []string
	for _, p := range strings.Split(c.Perks, ",") {
		perks = append(perks, strings.TrimSpace(p))
	}
	dto.Perks = perks

	u := c.User
	dto.User = campaign.DtoUserCampaignFormat{
		ID: u.ID,
		Name: u.Name,
		ImageUrl: u.AvatarFileName,
	}

	images:= []campaign.DtoImageCampaignFormat{}
	for _, i := range c.CampaignImages {
		data := campaign.DtoImageCampaignFormat{}
		isPrimary := false
		if i.IsPrimary == 1 {
			isPrimary = true
		}
		data.ImageUrl = i.FileName
		data.IsPrimary = isPrimary
		images = append(images, data)
	}
	dto.Images = images

	return dto
}
