package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID               uuid.UUID
	UserId           int
	CampaignId            uuid.UUID
	Amount       int
	Status    string
	Code             string
	PaymentUrl string  `gorm:"type:varchar(100)"`
	User             user.User
	Campaign campaign.Campaign
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.Must(uuid.NewRandom())
	return
}
