// Package rates:
package rates

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

const RatesFileName = "rates.json"

func UploadRates() (map[string]any, error) {
	var rates map[string]any
	file, err := os.Open(RatesFileName)
	if err != nil {
		return nil, errors.New("error opennig file: " + RatesFileName)
	}
	jsonString := ""
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		jsonString += scan.Text()
	}
	if err := json.Unmarshal([]byte(jsonString), &rates); err != nil {
		return nil, errors.New("error unmarshal " + RatesFileName)
	}

	return rates, nil
}
