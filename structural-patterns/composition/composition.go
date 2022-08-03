package composition

type Swimmer interface {
	Swim() string
}

type Athlete interface {
	Train() string
}

type SwimmerImpl struct {}

func (s *SwimmerImpl) Swim() string {
	return "Swimming"
}

type AthleteImpl struct {}

func (a *AthleteImpl) Train() string {
	return "Training"
}

type CompositeSwimmer struct {
	Athlete
	Swimmer
}

