package main

import (
	"fmt"
	"strings"
)

const (
	baseFare       = 50
	perMinuteRate  = 5
	economyRate    = 10
	standardRate   = 15
	businessRate   = 25
)

func calculateFare(carType string, distance float64, time int) (float64, string) {
	normalizedType := strings.ToLower(carType)
	var ratePerKm float64
	var description string

	switch normalizedType {
	case "эконом", "economy":
		ratePerKm = economyRate
		description = "Эконом класс"
	case "стандарт", "standard":
		ratePerKm = standardRate
		description = "Стандарт класс"
	case "бизнес", "business":
		ratePerKm = businessRate
		description = "Бизнес класс"
	}

	distanceCost := distance * ratePerKm
	timeCost := float64(time) * perMinuteRate
	totalFare := baseFare + distanceCost + timeCost

	return totalFare, description
}

func main() {
	fare1, desc1 := calculateFare("эконом", 10.5, 15)
	fmt.Printf("Тип: %s\nРасстояние: 10.5 км\nВремя: 15 мин\nСтоимость: %.2f руб.\n\n", desc1, fare1)

	fare2, desc2 := calculateFare("стандарт", 8.2, 10)
	fmt.Printf("Тип: %s\nРасстояние: 8.2 км\nВремя: 10 мин\nСтоимость: %.2f руб.\n\n", desc2, fare2)

	fare3, desc3 := calculateFare("бизнес", 5.0, 5)
	fmt.Printf("Тип: %s\nРасстояние: 5.0 км\nВремя: 5 мин\nСтоимость: %.2f руб.\n\n", desc3, fare3)
}