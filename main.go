package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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

func main() {
	apiKey := "ВАШ API-ключ" // API-ключ

	var city string
	fmt.Print("Введите город: ")
	fmt.Scanln(&city)

	city = url.QueryEscape(city) // подготовка к вставке в url

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=ru", city, apiKey)

	response, err := http.Get(url) // выполнение запроса
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK { // Проверка статуса ответа
		fmt.Printf("Ошибка: сервер вернул статус %d\n", response.StatusCode)
		return
	}

	body, err := ioutil.ReadAll(response.Body) // Чтение ответа
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	var weather WeatherResponse // Разбор JSON-ответа
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	if weather.Name == "" {
		fmt.Println("Город не найден. Проверьте правильность ввода.")
		return
	}

	fmt.Printf("\nПогода в городе %s:\n", weather.Name)
	fmt.Printf("Температура: %.1f°C\n", weather.Main.Temp)
	fmt.Printf("Ощущается как: %.1f°C\n", weather.Main.FeelsLike)
	fmt.Printf("Давление: %d hPa\n", weather.Main.Pressure)
	fmt.Printf("Влажность: %d%%\n", weather.Main.Humidity)
	fmt.Printf("Описание: %s\n", weather.Weather[0].Description)
	bufio.NewReader(os.Stdin).ReadBytes(' ')
}
