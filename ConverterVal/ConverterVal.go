package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/everapihq/freecurrencyapi-go"
)

func main() {
	// Инициализация API
	freecurrencyapi.Init("fca_live_fql8Ub9Hbh0CKe4iPmeWBAMUNR9FRn7IcPsUAZ4C")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите базовую валюту (например, USD): ")
	baseCurrency, _ := reader.ReadString('\n')
	baseCurrency = strings.TrimSpace(baseCurrency)

	fmt.Print("Введите целевую валюту (например, EUR): ")
	targetCurrency, _ := reader.ReadString('\n')
	targetCurrency = strings.TrimSpace(targetCurrency)

	fmt.Print("Введите сумму для конвертации: ")
	var amount float64
	_, err := fmt.Scanf("%f", &amount)
	if err != nil {
		log.Fatal(err)
	}

	// Создаем параметры для запроса курсов валют
	params := map[string]string{
		"base_currency": baseCurrency,
		"currencies":    targetCurrency,
	}

	// Получаем курсы валют
	latestRates := freecurrencyapi.Latest(params)

	// Парсим полученные данные
	var data map[string]interface{}
	if err := json.Unmarshal(latestRates, &data); err != nil {
		log.Fatal(err)
	}

	// Получаем курс конвертации
	rates := data["data"].(map[string]interface{})
	rate := rates[targetCurrency].(float64)

	// Выполняем конвертацию
	convertedAmount := amount * rate

	fmt.Printf("%.2f %s равны %.2f %s\n", amount, baseCurrency, convertedAmount, targetCurrency)
}
