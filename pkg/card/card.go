package card

type Transaction struct {
	Id     string
	Bill   int64
	Time   int64
	Mcc    string
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

func SumByMCC(transactions []Transaction, mcc []string) int64 {
	var mmcSum int64

	for _, code := range mcc {
		for _, t := range transactions {
			if code == t.Mcc {
				mmcSum += t.Bill
			}
		}
	}

	return mmcSum

}
