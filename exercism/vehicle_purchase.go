package main

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	msg := " is clearly the better choice."
	var choice string
	if option1 < option2 {
		choice = option1
	} else {
		choice = option2
	}
	return choice + msg
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var multiplier float64
	switch {
	case age < 3:
		multiplier = 80
	case age < 10:
		multiplier = 70
	default:
		multiplier = 50
	}
	return originalPrice * multiplier / 100
}
