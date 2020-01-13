package device

import (
    "encoding/json"
    "fmt"
    config "github.com/framled/simpleMDM/v1/config"
    "io/ioutil"
    "net/http"
    "strconv"
    "time"
)

type MDM struct {
    ApiKey string
    HttpClient *http.Client
    BaseURL string
}

func NewMDM(apiKey string, baseUrl string) Repository {
    if baseUrl == "" {
        baseUrl = config.BaseURL
    }
    return &MDM{
        ApiKey: apiKey,
        HttpClient: &http.Client{
            Timeout: time.Second,
        },
        BaseURL: baseUrl,
    }
}

func (mdm *MDM) findAll(limit int, startingAfter int) (*ResponseMany, error) {
    request, err := http.NewRequest("GET", fmt.Sprintf("%v/devices", mdm.BaseURL), nil)
    if err != nil {
        return nil, err
    }
    request.Header.Set("authorization", mdm.ApiKey)
    query := request.URL.Query()
    query.Add("limit", strconv.Itoa(limit))
    query.Add("starting_after", strconv.Itoa(startingAfter))
    request.URL.RawQuery = query.Encode()

    response, err := mdm.HttpClient.Do(request)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    deviceResponse := &ResponseMany{}
    if err := json.Unmarshal(body, deviceResponse); err != nil {
        return nil, err
    }

    return deviceResponse, nil
}

func (mdm *MDM) getById(deviceId string) (*ResponseOne, error) {
    request, err := http.NewRequest("GET", fmt.Sprintf("%v/devices/%v", mdm.BaseURL, deviceId), nil)
    if err != nil {
        return nil, err
    }
    request.Header.Set("authorization", mdm.ApiKey)
    response, err := mdm.HttpClient.Do(request)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    deviceResponse := &ResponseOne{}
    if err := json.Unmarshal(body, deviceResponse); err != nil {
        return nil, err
    }

    return deviceResponse, nil
}
