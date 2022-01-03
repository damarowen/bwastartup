package transaction

import (
	"bwastartup/campaign"
	_payment "bwastartup/payment"
	"errors"
	"github.com/google/uuid"
)

type ITransactionService interface {
	GetTransactionByCampaignId(id DtoTransactionByCampaignId) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input DtoCreateTransaction) (Transaction, error)
	//ProcessPayment(input TransactionNotificationInput) error
	//GetAllTransactions() ([]Transaction, error)
}

type TransactionService struct {
	TransactionRepository ITransactionRepository
	CampaignRepository    campaign.ICampaignRepository
	paymentService        _payment.Service
}

func NewTransactionService(
	TransactionRepo ITransactionRepository,
	CampaignRepo campaign.ICampaignRepository,
	paymentService _payment.Service,
) ITransactionService {
	return &TransactionService{TransactionRepo, CampaignRepo, paymentService}
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

func (s *TransactionService) CreateTransaction(input DtoCreateTransaction) (Transaction, error) {

	_, err := s.CampaignRepository.FindById(input.CampaignId)

	if err != nil {
		return Transaction{}, err
	}

	transaction := Transaction{
		CampaignId: input.CampaignId,
		Amount:     input.Amount,
		UserId:     input.User.ID,
		Status:     "pending",
	}
	newTransaction, err := s.TransactionRepository.SaveTransaction(transaction)
	if err != nil {
		return newTransaction, err
	}

	//midtrans logic
	paymentTransaction := _payment.Transaction{
		Id:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentUrl = paymentURL

	// update buat tambahin payment URL
	newTransaction, err = s.TransactionRepository.UpdateTransaction(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

//func (s *service) ProcessPayment(input TransactionNotificationInput) error {
//	transaction_id, _ := strconv.Atoi(input.OrderID)
//
//	transaction, err := s.repository.GetByID(transaction_id)
//	if err != nil {
//		return err
//	}
//
//	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
//		transaction.Status = "paid"
//	} else if input.TransactionStatus == "settlement" {
//		transaction.Status = "paid"
//	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
//		transaction.Status = "cancelled"
//	}
//
//	updatedTransaction, err := s.repository.Update(transaction)
//	if err != nil {
//		return err
//	}
//
//	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignID)
//	if err != nil {
//		return err
//	}
//
//	if updatedTransaction.Status == "paid" {
//		campaign.BackerCount = campaign.BackerCount + 1
//		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount
//
//		_, err := s.campaignRepository.Update(campaign)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func (s *service) GetAllTransactions() ([]Transaction, error) {
//	transactions, err := s.repository.FindAll()
//	if err != nil {
//		return transactions, err
//	}
//
//	return transactions, nil
//}
