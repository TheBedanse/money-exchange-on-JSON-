package get_xml

import (
	"context"
	"encoding/xml"
	"fmt"
	currency "github.com/vintrinsics/money-exchange/internal/model"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
	"net/http"
	"time"
)

func GetXMLExchange(ctx context.Context, date time.Time) (currency.ValCurs, error) {
	var Client = http.Client{Timeout: 15 * time.Second}
	Client.CloseIdleConnections()
	url := fmt.Sprintf("http://www.cbr.ru/scripts/XML_daily.asp?date_req=%s", date.Format("02/01/2006"))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return currency.ValCurs{}, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.63 Safari/537.36")

	resp, err := Client.Do(req)
	if err != nil {
		return currency.ValCurs{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Println("Response:", bodyString)

		err = fmt.Errorf("have status code %d (%s)", resp.StatusCode, resp.Status)
		log.Println(err.Error())
		return currency.ValCurs{}, err
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			err = fmt.Errorf("unknown charset: %s", charset)
			log.Println(err.Error())
			return nil, err

		}
	}

	var valCurs currency.ValCurs
	err = decoder.Decode(&valCurs)
	if err != nil {
		log.Println(err.Error())
		return currency.ValCurs{}, err
	}
	return valCurs, nil
}
