package abstractFactory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

const (
	SportMotorbikeType = 1
	CruiseMotorbikeType = 2
)

const (
	CarFactoryType = 1
	MotorbikeFactoryType = 2
)

type CarFactory struct{}
func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
	return new(LuxuryCar), nil
	case FamilyCarType:
	return new(FamilyCar), nil
	default:
	return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}

type MotorbikeFactory struct {}

func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
	return new(SportMotorbike), nil
	case CruiseMotorbikeType:
	return new(CruiseMotorbike), nil
	default:
	return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n",v))
	}
}

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
	return new(CarFactory), nil
	case MotorbikeFactoryType:
	return new(MotorbikeFactory), nil
	default:
	return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n",f))
	}
}