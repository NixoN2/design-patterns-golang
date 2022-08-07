package strategy

import "testing"

func TestStrategy_Print(t *testing.T) {
	var activeStrategy PrintStrategy
	activeStrategy = &ConsoleSquare{}
	err := activeStrategy.Print()
	if err != nil {
		t.Error(err)
	}
	activeStrategy = &ImageSquare{"./image.jpg"}
	err = activeStrategy.Print()
	if err != nil {
		t.Error(err)
	}
}