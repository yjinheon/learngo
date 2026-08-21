package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load() // .env
}

func fetchWeather(city string, apiKey string, ch chan<- string, wg *sync.WaitGroup) interface{} {
	var data struct {
		Main struct {
			Temp float64 `json: "temp"`
		} `json: "main"`
	}

	defer wg.Done()

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
		return data
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Printf("Error Decoding Weather Data for %s: %s\n", city, err)
		return data
	}
	ch <- fmt.Sprintf("This is the %s", city)

	return data
}

func main() {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	startNow := time.Now()

	cities := []string{"Toronto", "London", "Paris", "Tokyo"}

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, apiKey, ch, &wg)

	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Println("This Operation took ", time.Since(startNow))
}
