package main

func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime == 0 {
		preparationTime = 2
	}
	return preparationTime * len(layers)
}

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(amounts []float64, numPortions int) []float64 {
	result := make([]float64, len(amounts))
	for i := range amounts {
		result[i] = amounts[i] * float64(numPortions) / 2
	}
	return result
}
