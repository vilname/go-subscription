package webClient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"subscription-back/internal/model"
	"time"
)

type CartWebCitrus struct {
}

func NewCartWebCitrus() *CartWebCitrus {
	return &CartWebCitrus{}
}

func (CartWebclient *CartWebCitrus) Create(dto model.CreateCartDto) (model.WebCitrusSuccessResponse, error) {
	var webCitrusSuccessResponse model.WebCitrusSuccessResponse

	urlApi := os.Getenv("URL_WEB_CITRUS")
	clientId := os.Getenv("WC_CLIENT_ID")
	clientSecret := os.Getenv("WC_CLIENT_SECRET")

	var webCitrusModel model.WebCitrus
	webCitrusModel.ProjectId = 21
	webCitrusModel.Fields = dto

	jsonData, err := json.Marshal(webCitrusModel)

	if err != nil {
		return webCitrusSuccessResponse, err
	}

	req, err := http.NewRequest("POST", urlApi+"/passes", bytes.NewBuffer(jsonData))

	if err != nil {
		return webCitrusSuccessResponse, err
	}

	req.Header.Set("X-Client-Id", clientId)
	req.Header.Set("X-Client-Secret", clientSecret)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Do(req)

	if err != nil {
		return webCitrusSuccessResponse, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return webCitrusSuccessResponse, err
	}

	err = json.Unmarshal(body, &webCitrusSuccessResponse)
	if err != nil {
		return webCitrusSuccessResponse, err
	}

	return webCitrusSuccessResponse, nil
}
