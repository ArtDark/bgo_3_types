package card


type Transaction struct {
	Id string
	Bill int64
	Time int64
	Mcc string
	Status string


}

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

func AddTransaction(card *Card, transaction Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}



