package main

import (
	"encoding/json"
	"fmt"
)

// 이 예제는 OpenWeather API 응답 중 일부를 흉내 낸 JSON입니다.
//
// 중요한 구조:
// - "name"은 최상위 key입니다.
// - "main"도 최상위 key입니다.
// - "main"의 value는 다시 객체입니다.
// - "temp", "humidity", "pressure"는 "main" 객체 안의 key입니다.
const sampleWeatherJSON = `{
	"name": "Seoul",
	"main": {
		"temp": 293.15,
		"humidity": 62,
		"pressure": 1012
	},
	"wind": {
		"speed": 3.1
	}
}`

// WeatherData는 "이 JSON을 Go 코드에서 어떤 모양으로 다룰지"를 정한 타입입니다.
//
// 실제 프로젝트에서는 보통 이렇게 이름 있는 struct를 만듭니다.
// 이유:
// - fetchWeather 같은 함수의 반환 타입으로 쓰기 좋습니다.
// - 테스트에서 재사용하기 좋습니다.
// - data.Main.Temp처럼 코드가 읽기 쉬워집니다.
type WeatherData struct {
	// JSON의 "name"과 연결됩니다.
	Name string `json:"name"`

	// JSON의 "main"과 연결됩니다.
	//
	// Main이라는 Go 필드명은 꼭 Main이어야 하는 것은 아닙니다.
	// 하지만 JSON key가 "main"이고 날씨 API 문서에서도 main이라고 부르므로
	// 같은 이름을 쓰면 읽는 사람이 덜 헷갈립니다.
	Main MainWeatherData `json:"main"`

	// JSON의 "wind"와 연결됩니다.
	Wind WindData `json:"wind"`
}

// MainWeatherData는 JSON의 "main" 객체 안쪽 구조입니다.
//
// OpenWeather에서 "main"은 온도, 습도, 기압 같은 주요 날씨 수치를 묶은 객체입니다.
type MainWeatherData struct {
	Temp     float64 `json:"temp"`
	Humidity int     `json:"humidity"`
	Pressure int     `json:"pressure"`
}

// WindData는 JSON의 "wind" 객체 안쪽 구조입니다.
type WindData struct {
	Speed float64 `json:"speed"`
}

func main() {
	fmt.Println("case 1: anonymous struct")
	decodeWithAnonymousStruct()

	fmt.Println()
	fmt.Println("case 2: named struct")
	decodeWithNamedStructExample()

	fmt.Println()
	fmt.Println("case 3: Go field name can differ from JSON key")
	decodeWithDifferentGoFieldName()

	fmt.Println()
	fmt.Println("case 4: partial struct")
	decodeOnlyFieldsYouNeed()

	fmt.Println()
	fmt.Println("case 5: map[string]any")
	decodeWithMap()
}

func decodeWithAnonymousStruct() {
	// 경우 1: 익명 struct
	//
	// 언제 쓰나?
	// - JSON 구조가 간단합니다.
	// - 이 함수 안에서 딱 한 번만 씁니다.
	// - 타입 이름을 따로 만드는 것이 오히려 번거롭습니다.
	//
	// 단점:
	// - 다른 함수에서 같은 구조를 재사용하기 어렵습니다.
	// - 함수 반환 타입으로 쓰기 불편합니다.
	var data struct {
		Name string `json:"name"`

		// JSON "main"의 value가 객체이므로 Go에서도 nested struct로 받습니다.
		Main struct {
			Temp     float64 `json:"temp"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
	}

	if err := json.Unmarshal([]byte(sampleWeatherJSON), &data); err != nil {
		fmt.Println("decode error:", err)
		return
	}

	fmt.Println("city:", data.Name)
	fmt.Println("temperature:", data.Main.Temp)
	fmt.Println("humidity:", data.Main.Humidity)
}

func decodeWithNamedStruct(jsonText string) (WeatherData, error) {
	// 경우 2: 이름 있는 struct
	//
	// 언제 쓰나?
	// - 실제 프로젝트 코드입니다.
	// - 여러 함수에서 같은 응답 타입을 씁니다.
	// - fetchWeather 함수의 반환 타입으로 쓰고 싶습니다.
	// - 테스트에서 타입을 재사용하고 싶습니다.
	var data WeatherData
	err := json.Unmarshal([]byte(jsonText), &data)
	return data, err
}

func decodeWithNamedStructExample() {
	data, err := decodeWithNamedStruct(sampleWeatherJSON)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}

	// JSON "main.temp"는 Go에서 data.Main.Temp로 읽습니다.
	fmt.Println("city:", data.Name)
	fmt.Println("temperature:", data.Main.Temp)
	fmt.Println("humidity:", data.Main.Humidity)
	fmt.Println("wind speed:", data.Wind.Speed)
}

func decodeWithDifferentGoFieldName() {
	// 경우 3: Go 필드명과 JSON key 이름을 다르게 쓰기
	//
	// Main이라고 꼭 써야 하는 것은 아닙니다.
	// 아래에서는 Go 필드명을 WeatherNumbers라고 했지만,
	// struct tag가 json:"main"이므로 JSON의 "main" 값이 여기에 들어옵니다.
	//
	// 언제 쓰나?
	// - JSON key 이름이 Go 코드에서 의미가 애매합니다.
	// - Go 코드에서는 더 설명적인 이름을 쓰고 싶습니다.
	//
	// 주의:
	// - WeatherNumbers는 대문자로 시작해야 합니다.
	// - weatherNumbers처럼 소문자로 시작하면 encoding/json이 값을 채울 수 없습니다.
	var data struct {
		CityName string `json:"name"`

		WeatherNumbers struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	if err := json.Unmarshal([]byte(sampleWeatherJSON), &data); err != nil {
		fmt.Println("decode error:", err)
		return
	}

	fmt.Println("city:", data.CityName)
	fmt.Println("temperature:", data.WeatherNumbers.Temp)
}

func decodeOnlyFieldsYouNeed() {
	// 경우 4: 필요한 필드만 받기
	//
	// JSON에는 humidity, pressure, wind 같은 값이 있어도,
	// struct에 Temp만 정의하면 Temp만 채워지고 나머지는 무시됩니다.
	//
	// 언제 쓰나?
	// - API 응답은 크지만 지금 필요한 값은 일부입니다.
	// - 학습 단계에서 구조를 단순하게 보고 싶습니다.
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	if err := json.Unmarshal([]byte(sampleWeatherJSON), &data); err != nil {
		fmt.Println("decode error:", err)
		return
	}

	fmt.Println("temperature only:", data.Main.Temp)
}

func decodeWithMap() {
	// 경우 5: map[string]any
	//
	// struct를 미리 만들지 않고 JSON을 map으로 받을 수도 있습니다.
	//
	// 언제 쓰나?
	// - JSON 구조가 자주 바뀝니다.
	// - 어떤 key가 올지 미리 모릅니다.
	// - 임시로 응답 구조를 탐색하고 싶습니다.
	//
	// 단점:
	// - 타입 변환을 직접 해야 해서 코드가 지저분해집니다.
	// - main["temp"]가 진짜 숫자인지 매번 확인해야 합니다.
	// - 실제 프로젝트에서는 구조가 정해진 API라면 struct가 더 안전합니다.
	var data map[string]any
	if err := json.Unmarshal([]byte(sampleWeatherJSON), &data); err != nil {
		fmt.Println("decode error:", err)
		return
	}

	mainValue, ok := data["main"].(map[string]any)
	if !ok {
		fmt.Println(`"main" is missing or not an object`)
		return
	}

	temp, ok := mainValue["temp"].(float64)
	if !ok {
		fmt.Println(`"main.temp" is missing or not a number`)
		return
	}

	fmt.Println("temperature:", temp)
}
