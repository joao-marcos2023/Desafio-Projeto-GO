package main

import "fmt"

const boilingTemperatureK = 373

func main() {
	temperatureK := boilingTemperatureK
	temperatureC := temperatureK - 273

	fmt.Printf("Temperatura de ebulição da água em °Kelvin = %d e em °Celsius = %d", temperatureK, temperatureC)
}
