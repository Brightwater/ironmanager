package groupIron

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

type ApiClient struct {
	httpClient *http.Client
	baseURL    string
	token      string
}

// NewApiClient creates a new API client instance
func NewApiClient(baseURL string, token string) *ApiClient {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &ApiClient{
		httpClient: &http.Client{Transport: tr},
		baseURL:    baseURL,
		token:      token,
	}
}

func (client *ApiClient) GetData() (string, error) {
	req, err := http.NewRequest("GET", client.baseURL+"/get-group-data?from_time=1980-12-23T03:57:02.960Z", nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", client.token)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	fmt.Println(resp)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

//
func (client *ApiClient) GetXpAllTime() (string, error) {
	req, err := http.NewRequest("GET", client.baseURL+"/get-skill-data", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("period", "Year")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", client.token)

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	fmt.Println(resp)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}


