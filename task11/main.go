package main

import "fmt"

// Интерфейс
type PaymentProcessor interface {
	Process(amount float64) string
}

// Оплата банковской картой
type CreditCard struct {
	CardNumber string
	HolderName string
}

func (c CreditCard) Process(amount float64) string {
	return fmt.Sprintf(
		"Платеж %.2f руб. успешно выполнен с карты %s",
		amount,
		c.CardNumber,
	)
}

// Оплата криптокошельком
type CryptoWallet struct {
	WalletAddress string
	Currency      string
}

func (c CryptoWallet) Process(amount float64) string {
	return fmt.Sprintf(
		"Перевод %.2f %s успешно выполнен с кошелька %s",
		amount,
		c.Currency,
		c.WalletAddress,
	)
}

func main() {
	// Слайс интерфейсов
	payments := []PaymentProcessor{
		CreditCard{
			CardNumber: "**** **** **** 1234",
			HolderName: "Иван Иванов",
		},
		CryptoWallet{
			WalletAddress: "0xA1B2C3D4",
			Currency:      "BTC",
		},
	}

	// Обработка платежей
	for _, payment := range payments {
		fmt.Println(payment.Process(1000))
	}
}