package transaction

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	GetTransactionsByCampaignId(campaignId uuid.UUID) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)

}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{db}
}

func ( r *TransactionRepository) GetTransactionsByCampaignId(campaignId uuid.UUID) ([]Transaction, error) {
	var t []Transaction
	err := r.db.Preload("User").Where("campaign_id = ?", campaignId).Take(&t).Error
	if err != nil {
		return t, err
	}
	return t ,nil

}

func ( r *TransactionRepository) GetTransactionsByUserID(userId  int) ([]Transaction, error) {
	var t []Transaction

	//transaction berelasi dengan campaign
	// dari relasi tersebut kita ambil campaign image dari tabel campaign image
	err := r.db.Preload("Campaign.User").
		Preload("Campaign.CampaignImages", "campaign_images.is_primary = 1").
		Where("user_id = ?", userId).Order("created_at desc").Find(&t).Error
	if err != nil {
		return t, err
	}
	return t ,nil

}
