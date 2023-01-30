package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	//panic("CalculateWorkingCarsPerHour not implemented")
	var cars float64
	var percentRate = successRate/100
	cars = float64(productionRate) * percentRate
	return cars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	panic("CalculateWorkingCarsPerMinute not implemented")
	var cars float64
	var minRate = productionRate/60
	cars = float64(minRate) * successRate
	return int(cars)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	panic("CalculateCost not implemented")
}
