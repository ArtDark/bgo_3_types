package card_test

import (
	"fmt"
	"github.com/ArtDark/bgo_3_types/pkg/card"
)

func ExampleAddTransaction() {

	master := &card.Card{
		Id:       0001,
		Issuer:   "MasterCard",
		Balance:  43_422_43,
		Number:   "2720_9955_9454_5218",
		Currency: "RUB",
		Icon:     "http://...",
		Transactions: []card.Transaction{
			{
				Id:     "1",
				Bill:   1_203_91,
				Time:   1592594432,
				MCC:    "5812",
				Status: "Valid",
			},
			{
				Id:     "3",
				Bill:   735_55,
				Time:   1592667170,
				MCC:    "5411",
				Status: "Valid",
			},
			{
				Id:     "4",
				Bill:   455_99,
				Time:   1592842454,
				MCC:    "5931",
				Status: "Valid",
			},
		},
	}
	transaction := card.Transaction{
		Id:     "5",
		Bill:   233_43,
		Time:   1596773221,
		MCC:    "5411",
		Status: "Valid",
	}
	fmt.Println(*master)
	card.AddTransaction(master, transaction)
	fmt.Println(*master)

	// Output:
	// {1 MasterCard 4342243 RUB 2720_9955_9454_5218 http://... [{1 120391 1592594432 5812 Valid} {3 73555 1592667170 5411 Valid} {4 45599 1592842454 5931 Valid}]}
	// {1 MasterCard 4342243 RUB 2720_9955_9454_5218 http://... [{1 120391 1592594432 5812 Valid} {3 73555 1592667170 5411 Valid} {4 45599 1592842454 5931 Valid} {5 23343 1596773221 5411 Valid}]}

}
func ExampleSumByMCC() {
	master := &card.Card{
		Id:       0001,
		Issuer:   "MasterCard",
		Balance:  43_422_43,
		Number:   "2720_9955_9454_5218",
		Currency: "RUB",
		Icon:     "http://...",
		Transactions: []card.Transaction{
			{
				Id:     "1",
				Bill:   1_203_91,
				Time:   1592594432,
				MCC:    "5812",
				Status: "Valid",
			},
			{
				Id:     "3",
				Bill:   735_55,
				Time:   1592667170,
				MCC:    "5411",
				Status: "Valid",
			},
			{
				Id:     "4",
				Bill:   455_99,
				Time:   1592842454,
				MCC:    "5931",
				Status: "Valid",
			},
		},
	}
	mcc := []string{"5411", "5812"}
	cashBack := card.SumByMCC(master.Transactions, mcc)
	fmt.Println(cashBack)

	// Output:
	// 193946

}
func ExampleTranslateMCC() {
	mcc := "34534"
	fmt.Println(card.TranslateMCC(mcc))
	// Output:
	// Category undefined or absent

}
func ExampleLastNTransactions() {
	master := &card.Card{
		Id:       0001,
		Issuer:   "MasterCard",
		Balance:  43_422_43,
		Number:   "2720_9955_9454_5218",
		Currency: "RUB",
		Icon:     "http://...",
		Transactions: []card.Transaction{
			{
				Id:     "1",
				Bill:   1_203_91,
				Time:   1592594432,
				MCC:    "5812",
				Status: "Valid",
			},
			{
				Id:     "2",
				Bill:   1_203_91,
				Time:   1592594432,
				MCC:    "5812",
				Status: "Valid",
			},
			{
				Id:     "3",
				Bill:   735_55,
				Time:   1592667170,
				MCC:    "5411",
				Status: "Valid",
			},
			{
				Id:     "4",
				Bill:   455_99,
				Time:   1592842454,
				MCC:    "5931",
				Status: "Valid",
			},
			{
				Id:     "5",
				Bill:   233_43,
				Time:   1596773221,
				MCC:    "5411",
				Status: "Valid",
			},
		},
	}
	fmt.Println(card.LastNTransactions(master, 3))

	// Output:
	// [{5 23343 1596773221 5411 Valid} {4 45599 1592842454 5931 Valid} {3 73555 1592667170 5411 Valid}]
}
