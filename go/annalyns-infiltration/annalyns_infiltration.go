package annalyn

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	var fast bool
	if (knightIsAwake == true) {
		fast = false
	} else {
		fast = true
	}
	return fast
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	var spy bool
	if (knightIsAwake || archerIsAwake || prisonerIsAwake ) {
		spy = true
	} else {
		spy = false
	}
	return spy
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	var signal bool
	if !(archerIsAwake) && prisonerIsAwake {
		signal = true
	} else {
		signal = false
	}
	return signal
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	var free bool
	case1:= false
	case2:= false
	if (prisonerIsAwake && !knightIsAwake && !archerIsAwake) {
		case1 = true
	}
	if (petDogIsPresent && !archerIsAwake){
		case2 = true
	}
	if (case1 || case2) {
		free = true
	}

	return free
}
