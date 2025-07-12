package main

import (
	"fmt"
	"console-weather/internal/api"
	"console-weather/pkg/utils"
)

func main() {
	apiKey := "ВАШ_API_КЛЮЧ" // Замените на ваш API ключ

	city := utils.GetUserInput("Введите город: ")
	
	weather, err := api.GetWeather(city, apiKey)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		utils.WaitForEnter()
		return
	}

	weather.Display()
	utils.WaitForEnter()
}