package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var cars float64
	var percentRate = successRate/100
	cars = float64(productionRate) * percentRate
	return cars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var carsPerMin int
	var carsPerHour float64
	carsPerHour = CalculateWorkingCarsPerHour(productionRate, successRate)
	carsPerMin = int(carsPerHour)/60
	return carsPerMin
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var remainder int
	var costGroupsOfTen uint
	var costIndividualBuilds uint

	costGroupsOfTen = uint((carsCount/10) * 95000)
	remainder = carsCount%10
	costIndividualBuilds = uint(remainder*10000)
	return costGroupsOfTen + costIndividualBuilds
}
