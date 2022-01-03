package transaction

import (
	"bwastartup/user"
	"github.com/google/uuid"
	"time"
)


type DtoTransactionByCampaignId struct {
	// /:id buat akses itu
	// nanti pake c.bindingWithUri
	ID string  `uri:"id" binding:"required,uuid"`
	User user.User
}


type DtoResponseCampaignTransaction struct {
	ID        uuid.UUID       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}


type DtoMappingResponseUserTransactions struct {
	ID        uuid.UUID               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	ID        uuid.UUID               `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	User user.User
}