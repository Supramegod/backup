package speed

// TODO: define the 'Car' type struct
type car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int
}

// Car is a remote controlled car with a certain speed and battery drain.
type Car = car
// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

// TODO: define the 'Track' type struct
type track struct {
	distance int
}

// Track is a track with a certain distance.
type Track = track

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{	
		distance: distance,
	}
		
	panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// Calculate the number of times the car needs to drive to complete the track
	drivesNeeded := track.distance / car.speed
	if track.distance%car.speed != 0 {
		drivesNeeded++
	}

	// Calculate the total battery drain needed
	totalBatteryDrain := drivesNeeded * car.batteryDrain

	// Check if the car has enough battery
	return car.battery >= totalBatteryDrain
}
