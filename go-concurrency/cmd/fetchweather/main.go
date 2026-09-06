// fetchweather는 도시별 날씨를 순차적으로 조회하는 예제입니다.
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"

	"go-concurrency/internal/weather"
)

func init() {
	godotenv.Load() // .env
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
	for _, city := range cities {
		data, err := client.Fetch(context.Background(), city)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Printf("%s: %.2fK (humidity %d%%)\n", data.Name, data.Main.Temp, data.Main.Humidity)
	}
	fmt.Println("This Operation took", time.Since(start))
}
