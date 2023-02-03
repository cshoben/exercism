package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	license := false
	if kind == "car" {
		license = true
	}
	if kind == "truck" {
		license = true
	}
	return license
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var chosen string
	if option1 < option2 {
		chosen = option1
	} else {
		chosen = option2
	}
	return chosen + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var resell float64
	if age < 3 {
		resell = originalPrice * 0.80
	} else if age >= 10 {
		resell = originalPrice * 0.50
	} else {
		resell = originalPrice * 0.70
	}
	return resell
}
