package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {

	return float64(productionRate) * successRate / 100.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionPerMinute := float64(productionRate) / 60.0
	result := productionPerMinute * successRate / 100.0

	return int(result)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	absCarsCount := uint(carsCount)
	groups := absCarsCount / 10
	individuals := absCarsCount % 10
	result := groups*95000 + individuals*10000

	return result
}
