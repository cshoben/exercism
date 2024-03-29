package speed


type Car struct {
	battery int
	batteryDrain int
	speed int
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	Car := Car {
		battery: 100,
		batteryDrain: batteryDrain,
		speed: speed,
		distance: 0,
	}
	return Car
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	Track := Track {
		distance: distance,
	}
	return Track
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	newBattery := car.battery-car.batteryDrain
	switch {
	case newBattery >= 0:
		car.distance = car.distance + car.speed
		car.battery = newBattery
	case newBattery < 0:
		return car
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// drive the car while there is distance to drive.
	// if car battery does not go down, then it can not finish track, return false
	for track.distance > 0 {
		batteryStart := car.battery
		car = Drive(car)
		batteryEnd := car.battery
		if batteryStart == batteryEnd {
			return false
		}
		track.distance = track.distance - car.speed
	}

	return true

}
