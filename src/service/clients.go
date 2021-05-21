package service

import (
	"adinunno.fr/ubiquiti-influx-monitoring/src/infra"
	"adinunno.fr/ubiquiti-influx-monitoring/src/response"
	"encoding/json"
	"fmt"
)

func GetAvailabilities(config infra.Config) ([]response.AvailabilityResponse, error) {
	url := fmt.Sprintf("%s?segment=ap&startDate=2021-05-21&endDate=2022-05-31", config.Url)

	serverRequest, err := httpGET(url, config.UserAgent)

	if err != nil {
		serverRequest = nil
		return nil, err
	}

	var inter []response.AvailabilityResponse

	decoder := json.NewDecoder(serverRequest.Body)

	err = decoder.Decode(&inter)
	defer serverRequest.Body.Close()

	decoder = nil
	serverRequest = nil
	if err != nil {
		return nil, err
	}
	return inter, nil
}