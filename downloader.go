package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

const url = "http://www.cbr.ru/scripts/XML_daily.asp?date_req="

func downLoad(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Accept", "application/xml")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad response %d: %s", resp.StatusCode, body)
	}
	return io.ReadAll(resp.Body)
}

func generationDate() time.Time {

	now := time.Now()
	dateOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return dateOnly
}

func Collector(countDay int) (response [][]byte, err error) {
	response = make([][]byte, countDay)
	date := generationDate()

	for i := 0; i < countDay; i++ {
		raw, err := downLoad(fmt.Sprintf("%s%02d/%02d/%04d", url, date.Day(), date.Month(), date.Year()))
		if err != nil {
			return nil, err
		}

		response[i] = raw
		if err != nil {
			return nil, err
		}

		date = date.AddDate(0, 0, -1)
	}

	return response, nil
}
