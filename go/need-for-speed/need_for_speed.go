package speed

// TODO: define the 'Car' type struct
type Car struct {
	speed        int // speed of the car
	batteryDrain int // battery drain per drive
	battery      int // current battery level
	distance     int // distance driven
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
	//panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int // distance of the track
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance, // distance of the track
	}
	//panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	} else {
		car.distance += 0
	}
	return car
	//panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	if car.battery >= track.distance/car.speed*car.batteryDrain {
		return true
	} else {
		return false
	}
	//panic("Please implement the CanFinish function")
}
