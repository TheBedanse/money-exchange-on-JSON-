package currency

type CurrencyS struct {
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type ValCurs struct {
	Date  string   `xml:"Date,attr"`
	Name  string   `xml:"name,attr"`
	Items []Valute `xml:"Valute"`
}

type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  int32  `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  int32  `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}
type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}
