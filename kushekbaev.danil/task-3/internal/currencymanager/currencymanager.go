package currencymanager

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/net/html/charset"
)

const (
	filePermission      = 0o644
	directoryPermission = 0o755
)

func Read(path string) (data.Currencies, error) {
	var currenciesData data.Currencies

	file, err := os.Open(path)
	if err != nil {
		return currenciesData, err
	}

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	err = decoder.Decode(&currenciesData)
	if err != nil {
		return currenciesData, err
	}

	err = file.Close()
	if err != nil {
		return currenciesData, err
	}

	return currenciesData, nil
}

func Write(path string, currencies data.Currencies) error {
	data, err := json.MarshalIndent(currencies.AllCurrencies, "", "\t")
	if err != nil {
		return err
	}

	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, directoryPermission); err != nil {
		return err
	}

	err = os.WriteFile(path, data, filePermission)
	if err != nil {
		return err
	}

	return nil
}
