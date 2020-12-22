// Package for work with user cards
package card

import (
	"fmt"
)

// User transaction structure
type Transaction struct {
	Id     string
	Bill   int64
	Time   int64
	MCC    string
	Status string
}

// User card structure
type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

// Function add new transaction
func AddTransaction(card *Card, transaction Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}

// Function sum bill transactions by MMC
func SumByMCC(transactions []Transaction, mcc []string) int64 {
	var mmcSum int64

	for _, code := range mcc {
		for _, t := range transactions {
			if code == t.MCC {
				mmcSum += t.Bill
			}
		}
	}

	return mmcSum

}

// Function translate MMC code to human friendly language
func TranslateMCC(code string) string {
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5812": "Рестораны",
	}

	const ErrCategoryUndef = "Category undefined or absent"

	if value, ok := mcc[code]; ok {
		return value
	}

	return ErrCategoryUndef

}

// Function view last N transactions
func LastNTransactions(card *Card, n int) []Transaction {

	if len(card.Transactions) == 0 {
		fmt.Println("Transactions is absent")
		return nil
	}

	if n <= 0 {
		fmt.Println("Bad num")
		return nil
	}

	if n > len(card.Transactions) {
		fmt.Println("Num is over range transactions")
		return nil
	}

	t := card.Transactions // user transactions

	for i, j := 0, len(t)-1; i < j; i, j = i+1, j-1 { //revert user transactions
		t[i], t[j] = t[j], t[i]
	}

	return t[:n]
}
