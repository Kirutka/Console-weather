package models

import (
	"encoding/json"
	"fmt"
)

type WeatherResponse struct {
	Main struct {
		Temp      float64 `json:"temp"`       // Температура
		FeelsLike float64 `json:"feels_like"` // Ощущается как
		Pressure  int     `json:"pressure"`   // Атмосферное давление
		Humidity  int     `json:"humidity"`   // Влажность
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"` // Описание погоды
	} `json:"weather"`
	Name string `json:"name"` // Название города
}

// Display выводит информацию о погоде в консоль
func (w *WeatherResponse) Display() {
	if w.Name == "" {
		fmt.Println("Город не найден. Проверьте правильность ввода.")
		return
	}

	fmt.Printf("\nПогода в городе %s:\n", w.Name)
	fmt.Printf("Температура: %.1f°C\n", w.Main.Temp)
	fmt.Printf("Ощущается как: %.1f°C\n", w.Main.FeelsLike)
	fmt.Printf("Давление: %d hPa\n", w.Main.Pressure)
	fmt.Printf("Влажность: %d%%\n", w.Main.Humidity)
	
	if len(w.Weather) > 0 {
		fmt.Printf("Описание: %s\n", w.Weather[0].Description)
	} else {
		fmt.Println("Описание: недоступно")
	}
}

// UnmarshalJSON кастомный разбор JSON для обработки ошибок
func (w *WeatherResponse) UnmarshalJSON(data []byte) error {
	type Alias WeatherResponse
	aux := &struct {
		*Alias
		Cod int `json:"cod"`
	}{
		Alias: (*Alias)(w),
	}
	
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	
	return nil
}