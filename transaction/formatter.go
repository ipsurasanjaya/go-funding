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
