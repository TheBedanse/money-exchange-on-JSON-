package getExchange

import (
	"context"
	"encoding/json"
	"fmt"
	currency "github.com/vintrinsics/money-exchange/internal/model"
	getXML "github.com/vintrinsics/money-exchange/internal/service/get_xml"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetExchangeRate(currencyCode string) (float64, error) {
	ctx := context.Background()
	date := time.Now()
	valCurs, err := getXML.GetXMLExchange(ctx, date)
	if err != nil {
		return 0, err
	}

	rate := findExchangeRate(valCurs.Items, currencyCode)
	return rate, nil
}

func findExchangeRate(items []currency.Valute, currencyCode string) float64 {
	for _, item := range items {
		if item.CharCode == currencyCode {
			value := item.Value
			value = strings.Replace(value, ",", ".", -1)
			rate, _ := strconv.ParseFloat(value, 64)
			return rate
		}
	}
	return 0.0
}

func GetExchangeRates() (currency.ExchangeRates, error) {
	resp, err := http.Get("https://openexchangerates.org/api/latest.json?app_id=eb4cf019d8814dd78872b4143e76cc61")
	if err != nil {
		return currency.ExchangeRates{}, err
	}
	defer resp.Body.Close()

	var rates currency.ExchangeRates
	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		return currency.ExchangeRates{}, err
	}

	return rates, nil
}

func GetCurrencyRate(rates currency.ExchangeRates, currency string) (float64, error) {
	rate, ok := rates.Rates[currency]
	if !ok {
		return 0, fmt.Errorf("currency not found: %s", currency)
	}

	return rate, nil
}
