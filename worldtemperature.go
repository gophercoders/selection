package main

import (
	"fmt"

	"github.com/gophercoders/simpleio"
)

func main() {
	var temperature int

	var temperatureInReykjavik int
	var temperatureInAthens int
	var temperatureInBeijing int
	var temperatureInCaracas int
	var temperatureInHelsinki int
	var temperatureInLondon int
	var temperatureInMadrid int
	var temperatureInWashingtonDC int
	var temperatureInPretoria int
	var temperatureInSantiago int
	var temperatureInMoscow int
	var temperatureInTokyo int

	temperatureInReykjavik = 1
	temperatureInAthens = 17
	temperatureInBeijing = 17
	temperatureInCaracas = 32
	temperatureInHelsinki = 7
	temperatureInLondon = 11
	temperatureInMadrid = 13
	temperatureInWashingtonDC = 16
	temperatureInPretoria = 15
	temperatureInSantiago = 25
	temperatureInMoscow = 12
	temperatureInTokyo = 16

	fmt.Println("The worldtemperature program tells you which cities are ")
	fmt.Println("hotter or colder than where you live.")

	fmt.Println("Enter the temperature in degrees Celsius today: ")
	temperature = simpleio.ReadNumberFromKeyboard()

	fmt.Println("European Cities")
	if temperature > temperatureInReykjavik {
		fmt.Println("Hotter than Reykjavik in Iceland which is ")
		fmt.Print(temperatureInReykjavik)
		fmt.Println(" degrees Celsius")
	}
	if temperature > temperatureInHelsinki {
		fmt.Println("Hotter than Helsinki in Finland which is ")
		fmt.Print(temperatureInHelsinki)
		fmt.Println(" degrees Celsius")
	}
	if temperature <= temperatureInAthens {
		fmt.Println("Colder than Athens in Greece which is ")
		fmt.Print(temperatureInAthens)
		fmt.Println(" degrees Celsius")
	}
	if temperature == temperatureInLondon {
		fmt.Println("Exactly the same as London in the UK which is also ")
		fmt.Print(temperatureInLondon)
		fmt.Println(" degrees Celsius")
	}
	if temperature < temperatureInMadrid {
		fmt.Println("Colder than Madrid in Spain which is ")
		fmt.Print(temperatureInMadrid)
		fmt.Println(" degrees Celsius")
	}
	if temperature > temperatureInMoscow {
		fmt.Println("Hotter than Moscow in Russia which is ")
		fmt.Print(temperatureInMoscow)
		fmt.Println(" degrees Celsius")
	}
	fmt.Println("Asian Cities")
	if temperature > temperatureInBeijing {
		fmt.Println("Hotter than Beijing in China which is ")
		fmt.Print(temperatureInBeijing)
		fmt.Println(" degrees Celsius")
	}
	if temperature <= temperatureInTokyo {
		fmt.Println("Colder or the same as Tokyo in Japan which is ")
		fmt.Print(temperatureInTokyo)
		fmt.Println(" degrees Celsius")
	}
	fmt.Println("North American Cities")
	if temperature >= temperatureInWashingtonDC {
		fmt.Println("Warmer or the same as Washington DC in the USA which is ")
		fmt.Print(temperatureInWashingtonDC)
		fmt.Println(" degrees Celsius")
	}
	fmt.Println("Cities in Africa")
	if temperature > temperatureInPretoria {
		fmt.Println("Hotter than Pretoria in South Africa which is ")
		fmt.Print(temperatureInPretoria)
		fmt.Println(" degrees Celsius")
	}
	fmt.Println("Cities in South America")
	if temperature < temperatureInCaracas {
		fmt.Println("Colder than Caracas in Venezuela which is ")
		fmt.Print(temperatureInCaracas)
		fmt.Println(" degrees Celsius")
	}
	if temperature > temperatureInSantiago {
		fmt.Println("Hotter than Santiago in Chile which is ")
		fmt.Print(temperatureInSantiago)
		fmt.Println(" degrees Celsius")
	}
}
