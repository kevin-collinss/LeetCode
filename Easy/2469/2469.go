package Easy

func convertTemperature(celsius float64) []float64 {
	temp := make([]float64, 2)

	temp[0] = celsius + 273.15
	temp[1] = celsius*1.80 + 32.00

	return temp
}
