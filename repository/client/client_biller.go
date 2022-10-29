package client

import (
	"e-money-svc/config"
	"e-money-svc/domain/model"
	"e-money-svc/shared/enum"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type HttpClient struct {
	endpoint *url.URL
	header   map[string]string
	Client   *http.Client
}

func NewHttpClient() *HttpClient {
	cfg := config.GetConfig()
	url, _ := url.Parse(cfg.Client.Biller.Host)
	return &HttpClient{
		endpoint: url,
		header:   map[string]string{},
		Client:   &http.Client{},
	}
}

func (c *HttpClient) GetListBiller() (*model.ListBillerResponse, error) {
	var res model.ListBillerResponse
	url := fmt.Sprintf("%v%v", c.endpoint.Path, string(enum.BillerServiceGetList))
	req, err := http.NewRequest("GET", c.endpoint.Scheme+"://"+c.endpoint.Host+url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	fmt.Println()
	fmt.Println(res)
	if res.Code != 200 {
		return nil, errors.New(res.Status)
	}

	return &res, nil
}

func (c *HttpClient) GetDetailBiller(id string) (*model.DetailBillerResponse, error) {
	var res model.DetailBillerResponse
	c.endpoint.Path = path.Join(c.endpoint.Path, string(enum.BillerServiceDetail))

	params := url.Values{}
	params.Add("billerId", id)
	c.endpoint.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", c.endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	if res.Code != 200 {
		return nil, errors.New(res.Status)
	}

	return &res, nil
}
