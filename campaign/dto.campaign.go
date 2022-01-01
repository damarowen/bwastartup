package campaign

import (
	"bwastartup/user"
	"github.com/google/uuid"
	"time"
)

type DtoCampaign struct {
	ID             uuid.UUID `json:"id"`
	UserId           int `json:"user_id"`
	Name     string `json:"name"`
	ShortDescription          string `json:"short_description"`
	Slug  string`json:"slug"`
	ImageUrl   string `json:"image_url"`
	GoalAmount           int `json:"goal_amount"`
	CurrentAmount           int `json:"current_amount"`
	CreatedAt time.Time
	UpdatedAt 	time.Time
}


type DtoCampaignDetailById struct {
	// /:id buat akses itu
	// nanti pake c.bindingWithUri
	ID string  `uri:"id" binding:"required,uuid"`
}

type DtoCreateCampaign struct {
	Name     string `json:"name" binding:"required"`
	Description     string `json:"description" binding:"required"`
	ShortDescription          string `json:"short_description" binding:"required"`
	GoalAmount           int `json:"goal_amount" binding:"required"`
	Perks           string `json:"perks" binding:"required"`
	User user.User
}

type DtoUpdateCampaign struct {
	Name     string `json:"name" `
	Description     string `json:"description" `
	ShortDescription          string `json:"short_description" `
	GoalAmount           int `json:"goal_amount" `
	Perks           string `json:"perks" `
	User user.User
}

type DtoCampaignDetail struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"short_description"`
	Description      string    `json:"description"`
	ImageUrl         string    `json:"image_url"`
	GoalAmount       int       `json:"goal_amount"`
	CurrentAmount    int       `json:"current_amount"`
	Slug             string    `json:"slug"`
	Perks            []string  `json:"perks"`
	User DtoUserCampaignFormat `json:"user"`
	Images []DtoImageCampaignFormat `json:"images"`
}

type DtoUserCampaignFormat struct {
	ID              		int `json:"id"`
	Name             string    `json:"name"`
	ImageUrl         string    `json:"avatar_url"`
}

type DtoImageCampaignFormat struct {
	ImageUrl         string    `json:"image_url"`
	IsPrimary bool `json:"is_primary"`
}



