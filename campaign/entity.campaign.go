package campaign

import (
	"bwastartup/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Campaign struct {
	ID               uuid.UUID
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CampaignImages   []CampaignImage
	User             user.User
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CampaignImage struct {
	ID         int
	CampaignId uuid.UUID
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (c *Campaign) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.Must(uuid.NewRandom())

	//if !c.IsValid() {
	//	err = errors.New("can't save invalid data")
	//}
	return
}
