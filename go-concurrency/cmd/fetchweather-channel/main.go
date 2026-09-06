// fetchweather-channel은 goroutine + channel로 도시별 날씨를
// 동시에 조회하는 예제입니다.
package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"

	"go-concurrency/internal/weather"
)

func init() {
	godotenv.Load() // .env
}

type result struct {
	city string
	data weather.WeatherData
	err  error
}

func main() {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "OPENWEATHER_API_KEY is not set")
		os.Exit(1)
	}

	client := weather.NewClient(apiKey)
	cities := []string{"Toronto", "London", "Paris", "Tokyo"}

	start := time.Now()

	ch := make(chan result)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, err := client.Fetch(context.Background(), city)
			ch <- result{city: city, data: data, err: err}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		if r.err != nil {
			fmt.Fprintln(os.Stderr, r.err)
			continue
		}
		fmt.Printf("%s: %.2fK (humidity %d%%)\n", r.data.Name, r.data.Main.Temp, r.data.Main.Humidity)
	}

	fmt.Println("This Operation took", time.Since(start))
}
