package felidae

import "pprof/animal"

type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}
