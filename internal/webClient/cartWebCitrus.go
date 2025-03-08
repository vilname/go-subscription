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

func (cartWebclient *CartWebCitrus) Create(dto model.CreateCartDto) (model.WebCitrusSuccessResponse, error) {
	var webCitrusSuccessResponse model.WebCitrusSuccessResponse

	urlApi := os.Getenv("URL_WEB_CITRUS")

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

	webCitrusSuccessResponse, err = cartWebclient.responseClient(req)

	if err != nil {
		return webCitrusSuccessResponse, err
	}

	return webCitrusSuccessResponse, nil
}

func (cartWebclient *CartWebCitrus) GetElement(serialNumber string) (model.WebCitrusSuccessResponse, error) {
	var webCitrusSuccessResponse model.WebCitrusSuccessResponse

	urlApi := os.Getenv("URL_WEB_CITRUS")

	req, err := http.NewRequest("GET", urlApi+"/passes/"+serialNumber, nil)

	if err != nil {
		return webCitrusSuccessResponse, err
	}

	webCitrusSuccessResponse, err = cartWebclient.responseClient(req)

	if err != nil {
		return webCitrusSuccessResponse, err
	}

	return webCitrusSuccessResponse, nil
}

func (cartWebclient *CartWebCitrus) RemoveCardClient(serialNumber string) error {
	//passes/{serial_number}/registrations

	urlApi := os.Getenv("URL_WEB_CITRUS")

	req, err := http.NewRequest("DELETE", urlApi+"/passes/"+serialNumber+"/registrations", nil)

	if err != nil {
		return err
	}

	clientId := os.Getenv("WC_CLIENT_ID")
	clientSecret := os.Getenv("WC_CLIENT_SECRET")

	req.Header.Set("X-Client-Id", clientId)
	req.Header.Set("X-Client-Secret", clientSecret)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	_ = body

	if err != nil {
		return err
	}

	return nil
}

func (cartWebclient *CartWebCitrus) responseClient(req *http.Request) (model.WebCitrusSuccessResponse, error) {
	var webCitrusSuccessResponse model.WebCitrusSuccessResponse

	clientId := os.Getenv("WC_CLIENT_ID")
	clientSecret := os.Getenv("WC_CLIENT_SECRET")

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
