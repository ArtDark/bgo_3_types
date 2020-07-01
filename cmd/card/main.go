package main

import (
	"fmt"
	"github.com/ArtDark/bgo_3_types/pkg/card"
)

func main() {
	master := card.Card{
		Id:       0001,
		Issuer:   "MasterCard",
		Balance:  43_422_43,
		Number:   "2720_9955_9454_5218",
		Currency: "RUB",
		Icon:     "http://...",
		Transactions: []card.Transaction{
			card.Transaction{
				Id:     "1",
				Bill:   1_203_91,
				Time:   1592594432,
				Mcc:    "5812",
				Status: "Valid",
			},
			card.Transaction{
				Id:     "3",
				Bill:   735_55,
				Time:   1592667170,
				Mcc:    "5411",
				Status: "Valid",
			},
		},
	}
	masterPointer := &master

	transaction := card.Transaction{
		Id:     "5",
		Bill:   233_43,
		Time:   1596773221,
		Mcc:    "5411",
		Status: "Valid",
	}

	fmt.Println("MasterCard: ", master)
	card.AddTransaction(masterPointer, transaction)
	fmt.Println("MasterCard: ", master)
}
