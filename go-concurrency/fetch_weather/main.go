package main

import "fmt"
import "encoding/json"
import "net/http"
import "time"
import "github.com/joho/godotenv"
import "os"


func init() {
    godotenv.Load()  // .env
}

func fetchWeather(city string, apiKey string) interface {} {
	var data struct {
		Main struct {
			Temp float64 `json: "temp"`
		} `json: "main"`
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s",city, apiKey)

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n",city,err)
		return data
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Printf("Error Decoding Weather Data for %s: %s\n",city,err)
		return data
	}

	return data
	
}

func main() {
  apiKey := os.Getenv("OPENWEATHER_API_KEY")
	startNow := time.Now()

	cities := []string{"Toronto", "London","Paris","Tokyo"}

	for _ , city := range cities {
		data := fetchWeather(city, apiKey)

		fmt.Println("This is the data : ", data)
	}


	 fmt.Println("This Operation took ", time.Since(startNow))
	
}
