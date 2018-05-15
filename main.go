package main

import (
	"math"
	"os"
	"fmt"
	"github.com/aybabtme/uniplot/barchart"
)

func getPNull(lambda, mu float64) float64 {
	return 1-(lambda/mu)
}

func getProbability(lambda, mu, n float64) float64 {
	ro := lambda/mu
	return (math.Pow(ro,n)) * getPNull(lambda,  mu)
}

func getLengthOfSystem(lambda, mu float64) float64  {
	ro := lambda / mu
	return ro/(1-ro)
}

func getLengthOfQueue(lambda, mu float64) float64  {
	ro := lambda / mu
	return getLengthOfSystem(lambda, mu) - ro
}

func getExpectedWaitingTimeInQueue(lambda, mu float64) float64  {
	return getLengthOfQueue(lambda, mu) / lambda
}

func getExpectedWaitingTimeInSystem(lambda, mu float64) float64 {
	return getLengthOfSystem(lambda, mu) / lambda
}

func Round(val float64) float64 {
	return math.Round(val*100)/100
}

func main()  {
	var data [][2]int

	lambda := 1.0 // Rate
	mu := 1 / 0.6     // Server
	showProbabilities := false

	p0 := getPNull(lambda, mu)
	p1 := getProbability(lambda, mu, 1.0)
	pNoQueue := (p0 + getProbability(lambda, mu, 1.0))*100

	if showProbabilities {
		fmt.Println("P(n>=2) =", Round((1-p0-p1)*100), "%")
		fmt.Println("P(no queue) =", Round(pNoQueue), "%")
		fmt.Println("p( 0 ) =", Round(p0*100), "%")
	}

	p0Arr := [2]int{0, int(p0*1000)}
	data = append(data, p0Arr)

	for x := 1 ; x < 11 ; x++ {
		probability := getProbability(lambda, mu, float64(x))
		//p0Arr := [2]float64{0.0, p0}

		if showProbabilities {
			fmt.Println("p(", x, ") =", Round(probability*100),"%")
		}

		prob := [2]int{x, int(probability*1000)}
		data = append(data, prob)
	}

	fmt.Println("Expected length of System = Ls =", Round(getLengthOfSystem(lambda, mu)))
	fmt.Println("Expected length of Queue = Lq =", Round(getLengthOfQueue(lambda, mu)))

	waitingTimeInQueue := getExpectedWaitingTimeInQueue(lambda, mu)
	waitingTimeInSystem := getExpectedWaitingTimeInSystem(lambda, mu)

	fmt.Println("Expected waiting time in queue = ", waitingTimeInQueue,"\t=> Wq =", Round(waitingTimeInQueue*60), "minutes")
	fmt.Println("Expected waiting time in system = ", waitingTimeInSystem ,"\t=> Ws =", Round(waitingTimeInSystem*60), "minutes")

	plot := barchart.BarChartXYs(data)
	if err := barchart.Fprint(os.Stdout, plot, barchart.Linear(25)); err != nil {
		panic(err)
	}
}

