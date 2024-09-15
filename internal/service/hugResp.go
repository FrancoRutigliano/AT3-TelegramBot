package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HuggingFacePrediction struct {
	Label string  `json:"label"`
	Score float64 `json:"score"`
}

func IsNegative(message string, accessToken string) (bool, error) {

	url := "https://api-inference.huggingface.co/models/nlptown/bert-base-multilingual-uncased-sentiment"

	payload := map[string]string{
		"inputs": message,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return false, errors.New("error al serializar el payload")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return false, errors.New("error al crear la petición HTTP")
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, errors.New("error al realizar la petición a Hugging Face")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("respuesta no exitosa de Hugging Face: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.New("error al leer el body de la respuesta")
	}

	var result [][]HuggingFacePrediction
	if err = json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error al parsear la respuesta JSON:", err) // Log del error
		return false, errors.New("error al parsear la respuesta JSON de Hugging Face")
	}

	if result[0][0].Label == "1 star" {
		return true, nil
	}

	return false, nil
}
