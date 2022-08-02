package transaction

import "time"

type FormatTransaction struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func TransactionFormatter(transaction Transaction) FormatTransaction {
	var formatTransaction FormatTransaction
	formatTransaction.ID = transaction.ID
	formatTransaction.Name = transaction.User.Name
	formatTransaction.Amount = transaction.Amount
	formatTransaction.CreatedAt = transaction.CreatedAt

	return formatTransaction
}

func TransactionsFormatter(transactions []Transaction) []FormatTransaction {
	if len(transactions) == 0 {
		return []FormatTransaction{}
	}

	var formattedTransactions []FormatTransaction

	for _, transaction := range transactions {
		formattedTransaction := TransactionFormatter(transaction)
		formattedTransactions = append(formattedTransactions, formattedTransaction)
	}

	return formattedTransactions
}

type FormatUserTransaction struct {
	ID        int                       `json:"id"`
	Amount    int                       `json:"amount"`
	Status    string                    `json:"status"`
	CreatedAt time.Time                 `json:"created_at"`
	Campaign  FormatTransactionCampaign `json:"campaign"`
}

type FormatTransactionCampaign struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TransactionCampaignFormatter(transaction Transaction) FormatUserTransaction {
	var userTransaction FormatUserTransaction

	userTransaction.ID = transaction.ID
	userTransaction.Amount = transaction.Amount
	userTransaction.Status = transaction.Status
	userTransaction.CreatedAt = transaction.CreatedAt
	userTransaction.Campaign.Name = transaction.Campaign.Name
	userTransaction.Campaign.ImageUrl = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		userTransaction.Campaign.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}

	return userTransaction
}

func UserTransactionsFormatter(transactions []Transaction) []FormatUserTransaction {
	var userTransactions []FormatUserTransaction

	if len(transactions) == 0 {
		return []FormatUserTransaction{}
	}

	for _, transaction := range transactions {

		userTransactions = append(userTransactions, TransactionCampaignFormatter(transaction))
	}

	return userTransactions
}
