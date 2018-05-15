package Queue

import (
	"fmt"
	"math"
)

type Queue struct {
	Ro ro
}

func (queue Queue) Init(lambda, mu float64) Queue {
	return Queue{Ro: ro{lambda, mu}}
}

type ro struct {
	Lambda float64
	Mu     float64
}

// ro = lambda / mu
func (ro ro) calculated() float64 {
	return ro.Lambda / ro.Mu
}

func (queue Queue) PrintProbabilities() {
	p0 := queue.PNull()
	p1 := queue.Probability(1.0)
	pNoQueue := (p0 + queue.Probability(1.0)) * 100

	fmt.Printf("P(n>=2) = %v%% \n", Round((1-p0-p1)*100))
	fmt.Printf("P(no Queue) = %v%% \n", Round(pNoQueue))
	fmt.Printf("p(0) = %v%% \n", Round(p0*100))

	for x := 1; x < 11; x++ {
		pn := queue.Probability(float64(x))
		fmt.Printf("p(%v) = %v%% \n", x, Round(pn*100))
	}

	fmt.Printf("\n")
}

func (queue Queue) PNull() float64 {
	return 1 - queue.Ro.calculated()
}

func (queue Queue) Probability(n float64) float64 {
	return (math.Pow(queue.Ro.calculated(), n)) * queue.PNull()
}

func (queue Queue) LengthOfSystem() float64 {
	return queue.Ro.calculated() / (1 - queue.Ro.calculated())
}

func (queue Queue) LengthOfQueue() float64 {
	return queue.LengthOfSystem() - queue.Ro.calculated()
}

func (queue Queue) WaitingTimeInQueue() float64 {
	return queue.LengthOfQueue() / queue.Ro.Lambda
}

func (queue Queue) WaitingTimeInSystem() float64 {
	return queue.LengthOfSystem() / queue.Ro.Lambda
}

func Round(val float64) float64 {
	return math.Round(val*100) / 100
}
