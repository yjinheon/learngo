package main

import "testing"

func TestDecodeWithNamedStruct(t *testing.T) {
	weather, err := decodeWithNamedStruct(sampleWeatherJSON)
	if err != nil {
		t.Fatalf("decodeWithNamedStruct returned error: %v", err)
	}

	if weather.Name != "Seoul" {
		t.Fatalf("Name = %q, want Seoul", weather.Name)
	}

	if weather.Main.Temp != 293.15 {
		t.Fatalf("Main.Temp = %v, want 293.15", weather.Main.Temp)
	}

	if weather.Main.Humidity != 62 {
		t.Fatalf("Main.Humidity = %v, want 62", weather.Main.Humidity)
	}
}
