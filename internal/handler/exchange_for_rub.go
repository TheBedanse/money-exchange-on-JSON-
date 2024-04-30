package exchangeHandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	getExchange "github.com/vintrinsics/money-exchange/internal/service"
)

func ExchangeHandler() *gin.Engine {
	router := gin.Default()

	router.GET("/exchange-for-rub", func(c *gin.Context) {
		rateUSD, err := getExchange.GetExchangeRate("USD")
		rateEUR, err := getExchange.GetExchangeRate("EUR")
		rateAMD, err := getExchange.GetExchangeRate("AMD")
		resultAMD := rateAMD / 100
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get exchange rate"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"ЦБР-Доллар к рублю": rateUSD, "ЦБР-Евро к рублю": rateEUR, "ЦБР-Драм к рублю": resultAMD})
	}, func(c *gin.Context) {
		rates, err := getExchange.GetExchangeRates()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		usdRub, err := getExchange.GetCurrencyRate(rates, "RUB")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		eurUsd, err := getExchange.GetCurrencyRate(rates, "EUR")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		amdUsd, err := getExchange.GetCurrencyRate(rates, "AMD")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		amdRub := usdRub / amdUsd
		eurRub := usdRub / eurUsd
		formatAmdRub := fmt.Sprintf("%.3f", amdRub)
		formatEurRub := fmt.Sprintf("%.3f", eurRub)
		formatUsdRub := fmt.Sprintf("%.3f", usdRub)

		formatParseAmdRub, err := strconv.ParseFloat(formatAmdRub, 64)
		formatParseEurRub, err := strconv.ParseFloat(formatEurRub, 64)
		formatParseUsdRub, err := strconv.ParseFloat(formatUsdRub, 64)
		response := map[string]float64{
			"Openexchange-Доллар к рублю": formatParseUsdRub,
			"Openexchange-Евро к рублю":   formatParseEurRub,
			"Openexchange-Драм к рублю":   formatParseAmdRub,
		}

		c.JSON(http.StatusOK, response)
	})

	return router
}
