package chinese_holidays_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	UrlPrefix = "http://chinese-holidays-data.basten.me/data"
)

func loadData(urlPrefix string) ([]event, error) {
	url := fmt.Sprintf("%s/index.json", urlPrefix)
	bytes, err := downloadData(url)
	if err != nil {
		return nil, err
	}
	return newEvents(urlPrefix, bytes)
}

func downloadData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("下载数据错误，错误代码%d", resp.StatusCode))
	}
	defer func() { _ = resp.Body.Close() }()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

type entry struct {
	Year         uint   `json:"year"`
	LastModified string `json:"last_modified"`
}

func newEvents(urlPrefix string, data []byte) ([]event, error) {
	entries := make([]entry, 0)
	err := json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}

	events := make([]event, 0)
	for _, entry := range entries {
		url := fmt.Sprintf("%s/%d.json", urlPrefix, entry.Year)
		bytes, err := downloadData(url)
		if err != nil {
			return nil, err
		}

		e, err := parseEvents(bytes)
		if err != nil {
			return nil, err
		}
		events = append(events, e...)
	}
	return events, nil
}

func parseEvents(b []byte) ([]event, error) {
	var events []event
	err := json.Unmarshal(b, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}
