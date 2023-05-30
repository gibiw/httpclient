package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseUrl string
	token   string
	client  *http.Client
}

func New(baseUrl, token string, timeout time.Duration) *Client {
	return &Client{
		baseUrl: baseUrl,
		token:   fmt.Sprintf("PrivateToken %s", token),
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) GetParameterById(id string) (ParameterResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v2/parameters/%s", c.baseUrl, id), nil)
	if err != nil {
		return ParameterResponse{}, err
	}

	req.Header.Add("Authorization", c.token)
	req.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(req)
	if err != nil {
		return ParameterResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ParameterResponse{}, err
	}

	var parameter ParameterResponse
	err = json.Unmarshal(body, &parameter)
	if err != nil {
		return ParameterResponse{}, err
	}

	return parameter, nil
}

func (c *Client) AddParameter(name, value string) (ParameterResponse, error) {
	pr := ParameterRequest{
		Value: value,
		Name:  name,
	}
	data, err := json.Marshal(pr)
	if err != nil {
		return ParameterResponse{}, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v2/parameters", c.baseUrl), bytes.NewBuffer(data))
	if err != nil {
		return ParameterResponse{}, err
	}
	req.Header.Add("Authorization", c.token)
	req.Header.Add("Content-Type", "application/json")

	response, err := c.client.Do(req)
	if err != nil {
		return ParameterResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return ParameterResponse{}, err
	}

	var parameter ParameterResponse
	err = json.Unmarshal(body, &parameter)
	if err != nil {
		return ParameterResponse{}, err
	}

	return parameter, nil
}
