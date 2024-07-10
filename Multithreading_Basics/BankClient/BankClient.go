package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankClient interface {
	Deposit(amount int)
	Withdrawal(amount int) error
	Balance() int
}

type BankClientCell struct {
	balance int
}

func (b *BankClientCell) Deposit(amount int) {
	b.balance += amount
}

func (b *BankClientCell) Balance() int {
	return b.balance
}

func (b *BankClientCell) Withdrawal(amount int) error {
	if b.balance < amount {
		return errors.New("There are not enough funds in the account")
	}
	b.balance -= amount
	return nil
}

func NewBankClientCell(balance int) *BankClientCell {
	return &BankClientCell{
		balance: balance,
	}
}
func depositRandomAmount(bankClient *BankClientCell, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	amount := rand.Intn(10) + 1
	bankClient.Deposit(amount)
	fmt.Println("Deposited amount by goroutinne: $", amount)
	time.Sleep(time.Duration(rand.Float64()*0.5+0.5) * time.Second)
	mu.Unlock()
	wg.Done()
}
func withdrawalRandomAmount(bankClient *BankClientCell, wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	amount := rand.Intn(5) + 1
	bankClient.Withdrawal(amount)
	fmt.Println("Withdrawal amount by goroutine: $", amount)
	time.Sleep(time.Duration(rand.Float64()*0.5+0.5) * time.Second)
	mu.Unlock()
	wg.Done()
}

func main() {
	var baseBalance int
	fmt.Println("Enter initial balance:")
	fmt.Scan(&baseBalance)
	bankClient := NewBankClientCell(baseBalance)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go depositRandomAmount(bankClient, &wg, &mu)
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go withdrawalRandomAmount(bankClient, &wg, &mu)
	}
	for {
		var action string
		fmt.Println("Enter action (balance/deposit/withdrawal/exit):")
		fmt.Scan(&action)

		switch action {
		case "balance":
			fmt.Println("Balance in the account: $", bankClient.Balance())
		case "deposit":
			var depositAmount int
			fmt.Println("Enter deposit amount:")
			fmt.Scan(&depositAmount)
			bankClient.Deposit(depositAmount)
			fmt.Println("Deposited amount: $", depositAmount)
		case "withdrawal":
			var withdrawalAmount int
			fmt.Println("Enter withdrawal amount:")
			fmt.Scan(&withdrawalAmount)
			err := bankClient.Withdrawal(withdrawalAmount)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Withdrawn amount: $", withdrawalAmount)
			}
		case "exit":
			wg.Wait()
			return
		default:
			fmt.Println("Invalid action. Please try again.")
		}
	}
}
