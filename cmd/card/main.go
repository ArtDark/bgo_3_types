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
				MCC:    "5812",
				Status: "Valid",
			},
			card.Transaction{
				Id:     "3",
				Bill:   735_55,
				Time:   1592667170,
				MCC:    "5411",
				Status: "Valid",
			},
			card.Transaction{
				Id:     "4",
				Bill:   455_99,
				Time:   1592842454,
				MCC:    "5931",
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

	mcc := []string{"5411", "5812"}

	cashBack := card.SumByMCC(master.Transactions, mcc)
	fmt.Println("Cashback sum:", cashBack)


	category := card.TranslateMCC(master.Transactions[0].MCC)
	fmt.Println(category)

}
