package client

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type AgifyResponse struct {
    Age int `json:"age"`
}

type GenderizeResponse struct {
    Gender string `json:"gender"`
}

type NationalizeResponse struct {
    Country []struct {
        CountryID string `json:"country_id"`
    } `json:"country"`
}

func GetAge(name string) (int, error) {
    url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    var agify AgifyResponse
    if err := json.NewDecoder(resp.Body).Decode(&agify); err != nil {
        return 0, err
    }

    return agify.Age, nil
}

func GetGender(name string) (string, error) {
    url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var genderize GenderizeResponse
    if err := json.NewDecoder(resp.Body).Decode(&genderize); err != nil {
        return "", err
    }

    return genderize.Gender, nil
}

func GetNationality(name string) (string, error) {
    url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var nationalize NationalizeResponse
    if err := json.NewDecoder(resp.Body).Decode(&nationalize); err != nil {
        return "", err
    }

    if len(nationalize.Country) > 0 {
        return nationalize.Country[0].CountryID, nil
    }

    return "", nil
}
