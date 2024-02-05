package lasagna

func PreparationTime(layers []string, avePrepTime int) int {
	if avePrepTime == 0 {
		avePrepTime = 2
	}

	return len(layers) * avePrepTime
}

func Quantities(layers []string) (gramsOfNoodles int, litersOfSauce float64) {
	gramsOfNoodles = 0
	for _, v := range layers {
		if v == "noodles" {
			gramsOfNoodles += 50
		}
		if v == "sauce" {
			litersOfSauce += 0.2
		}
	}

	return
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(twoPortionAmts []float64, portionsToCook int) (scaledRecipe []float64) {
	scaledRecipe = append(make([]float64, 0, len(twoPortionAmts)), twoPortionAmts...)
	for i, v := range scaledRecipe {
		scaledRecipe[i] = (v / 2) * float64(portionsToCook)
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
