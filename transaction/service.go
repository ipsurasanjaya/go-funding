package transaction

import (
	"crowdfunding/campaign"
	"crowdfunding/payment"
	"errors"
	"fmt"

	"github.com/thanhpk/randstr"
)

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{
		repository:         repository,
		campaignRepository: campaignRepository,
		paymentService:     paymentService,
	}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if input.User.Id != campaign.UserID {
		return []Transaction{}, errors.New("not an owner of this campaign")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	var (
		transaction        Transaction
		paymentTransaction payment.Transaction
	)

	transaction.UserID = input.User.Id
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.TrxCode = GenerateTransactionCode()
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction.Amount = newTransaction.Amount
	paymentTransaction.ID = newTransaction.ID

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func GenerateTransactionCode() string {
	const codeLength = 15
	var (
		trxCode    string
		randString = randstr.String(codeLength, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	)

	trxCode = fmt.Sprintf("GFND-%s", randString)

	return trxCode
}
