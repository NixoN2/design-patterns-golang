package composition

import "testing"

func TestComposition(t *testing.T) {
	swimmer := CompositeSwimmer{
		&AthleteImpl{},
		&SwimmerImpl{},
	}

	if swimmer.Swim() != "Swimming" {
		t.Fatal("Swimmer can't swim")
	}

	if swimmer.Train() != "Training" {
		t.Fatal("Swimmer can't train")
	}

	t.Log(swimmer.Swim())
	t.Log(swimmer.Train())
}	