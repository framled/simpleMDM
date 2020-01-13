package device

import (
    "net/http"
)

const (
    BaseURL = "https://a.simplemdm.com/api/v1"
)

type MDM struct {
    ApiKey string
    HttpClient *http.Client
    BaseUrl string
}
