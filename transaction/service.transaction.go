package transaction

import (
	"bwastartup/campaign"
	"github.com/google/uuid"
	"errors"
)


type ITransactionService interface {
	GetTransactionByCampaignId(id DtoTransactionByCampaignId) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)

}

type TransactionService struct {
	TransactionRepository ITransactionRepository
	CampaignRepository campaign.ICampaignRepository
}

func NewTransactionService(TransactionRepo ITransactionRepository, CampaignRepo campaign.ICampaignRepository) ITransactionService {
	return &TransactionService{TransactionRepo, CampaignRepo}
}


func (s *TransactionService) GetTransactionByCampaignId(input DtoTransactionByCampaignId) ([]Transaction, error) {


	_ID, _ := uuid.Parse(input.ID)

	c, err := s.CampaignRepository.FindById(_ID)

	if c.UserId != input.User.ID {
		return []Transaction{}, errors.New("not authorize, different user")
	}

	if err != nil {
		return []Transaction{}, err
	}

	t, errs := s.TransactionRepository.GetTransactionsByCampaignId(_ID)

	if errs != nil {
		return t, err
	}
	return t, nil
}

func (s *TransactionService) GetTransactionsByUserID(userID int) ([]Transaction, error) {

	transactions, err := s.TransactionRepository.GetTransactionsByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
