package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"console-weather/internal/models"
)

// GetWeather получает данные о погоде для указанного города
func GetWeather(city, apiKey string) (*models.WeatherResponse, error) {
	encodedCity := url.QueryEscape(city)
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=ru", 
		encodedCity, apiKey)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("сервер вернул статус %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении ответа: %w", err)
	}

	var weather models.WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("ошибка при разборе JSON: %w", err)
	}

	return &weather, nil
}