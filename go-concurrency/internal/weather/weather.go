// Package weather는 OpenWeatherMap API에서 날씨 데이터를 가져오는
// 작은 클라이언트를 제공합니다.
package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const apiBaseURL = "https://api.openweathermap.org/data/2.5/weather"

// WeatherData는 OpenWeatherMap 응답 중 필요한 부분만 담는 타입입니다.
type WeatherData struct {
	Name string          `json:"name"`
	Main MainWeatherData `json:"main"`
	Wind WindData        `json:"wind"`
}

// MainWeatherData는 온도, 습도, 기압 같은 주요 날씨 수치를 담습니다.
type MainWeatherData struct {
	Temp     float64 `json:"temp"`
	Humidity int     `json:"humidity"`
	Pressure int     `json:"pressure"`
}

// WindData는 바람 관련 수치를 담습니다.
type WindData struct {
	Speed float64 `json:"speed"`
}

// Client는 OpenWeatherMap API 클라이언트입니다.
type Client struct {
	apiKey string
	http   *http.Client
}

// NewClient는 apiKey를 사용하는 Client를 만듭니다.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		http:   &http.Client{Timeout: 10 * time.Second},
	}
}

// Fetch는 도시의 현재 날씨를 조회합니다.
func (c *Client) Fetch(ctx context.Context, city string) (WeatherData, error) {
	var data WeatherData

	url := fmt.Sprintf("%s?q=%s&appid=%s", apiBaseURL, city, c.apiKey)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return data, fmt.Errorf("weather: building request for %s: %w", city, err)
	}

	res, err := c.http.Do(req)
	if err != nil {
		return data, fmt.Errorf("weather: fetching %s: %w", city, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("weather: %s: unexpected status %s", city, res.Status)
	}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("weather: decoding response for %s: %w", city, err)
	}

	return data, nil
}
