package main

import (
	"fmt"
	"github.com/ArtDark/bgo_3_types/pkg/card"
)

func main() {
	master := &card.Card{
		Id: 0001,
		Issuer: "MasterCard",
		Balance: 43_422_43,
		Number: "2720_9955_9454_5218",
		Currency: "RUB",
		Icon: "http://...",
		Transactions: []card.Transaction{
			card.Transaction{
				Id: 1,
				Bill: -1_203_91,
				Time: 1592594432,
				Mcc: "5812",
				Status: "Valid",
			},
			{
				Id:     3,
				Bill:   -735_55,
				Time:   1592667170,
				Mcc:    "5411",
				Status: "Valid",
			},
		},// TODO: слайс с двумя или тремя элементами (нужно отобразить те, что на скриншотах - время выберите сами)
	}
	fmt.Println(master)

}
