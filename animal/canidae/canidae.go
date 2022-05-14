package canidae

import "pprof/animal"

type Canidae interface {
	animal.Animal
	Run()
	Howl()
}
