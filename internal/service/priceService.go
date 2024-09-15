package service

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type DexToolsResp struct {
	Price float64 `json:"price"`
}

func GetTokenPrice() (float64, error) {
	url := "https://api.dextools.io/v2/token/polygon/0x2e8f3b0e4ad32317f70f7f79a63a1538ded23fd4/price"

	// falta access token desde dextools!!!!!!!!!!!

	resp, err := http.Get(url)
	if err != nil {
		return 0, errors.New("error al realizar la petición")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("respuesta no exitosa")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.New("error al leer el body de la respuesta")
	}

	var payload DexToolsResp

	if err = json.Unmarshal(body, &payload); err != nil {
		return 0, errors.New("error al parsear el JSON")
	}

	return payload.Price, nil
}
