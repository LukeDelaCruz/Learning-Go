package lasagna

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// Takes the number of layers you added to the lasagna as a parameter and returns how many minutes you spent preparing
// the lasagna, assuming each layer takes you 2 minutes to prepare.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna.
// This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
