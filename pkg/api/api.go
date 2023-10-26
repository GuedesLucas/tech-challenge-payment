package api

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

type APIClient struct {
	BaseURL string
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{BaseURL: baseURL}
}

func (client *APIClient) MakeRequest(method string, path string, headers map[string]string, queryParams map[string]string, body []byte) ([]byte, error) {
	// Construa o URL com base no caminho e nos parâmetros
	urlValues := url.Values{}
	for key, value := range queryParams {
		urlValues.Add(key, value)
	}
	url := client.BaseURL + path + "?" + urlValues.Encode()

	// Crie a solicitação HTTP com o método, URL e corpo (se houver)
	var requestBody *bytes.Reader
	if len(body) > 0 {
		requestBody = bytes.NewReader(body)
	}
	request, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, err
	}

	// Adicione os cabeçalhos à solicitação
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	// Realize a chamada HTTP
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Leia o corpo da resposta
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
