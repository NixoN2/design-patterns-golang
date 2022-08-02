package abstractFactory

type Motorbike interface {
	GetMotorbikeType() int
}

type SportMotorbike struct {}

func (s *SportMotorbike) NumWheels() int {
	return 2
}

func (s *SportMotorbike) NumSeats() int {
	return 1
}

func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}

type CruiseMotorbike struct {}

func (s *CruiseMotorbike) NumWheels() int {
	return 2
}

func (s *CruiseMotorbike) NumSeats() int {
	return 2
}

func (s *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}