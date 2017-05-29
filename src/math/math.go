package math

/**
 * Average
 * Calculates the mean average for provided array of numbers
 * @type {[type]}
 */
func Average(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	var total float64

	for _, n := range numbers {
		total += n
	}

	return total / float64(len(numbers))
}
