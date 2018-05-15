package main

import (
	"fmt"
	"github.com/aybabtme/uniplot/barchart"
	"m-m-1-ininity-go/Queue"
	"os"
)

func main() {
	var queue Queue.Queue

	queue = queue.Init(1.0, 1/0.6)

	if showProbabilities := true; showProbabilities {
		queue.PrintProbabilities()
	}

	var data [][2]int

	fmt.Printf("given values: \n"+
		"lamda = %v \n"+
		"mu = %v \n\n", queue.Ro.Lambda, queue.Ro.Mu)

	p0Arr := [2]int{0, int(queue.PNull() * 1000)}
	data = append(data, p0Arr)

	for x := 1; x < 11; x++ {
		prob := [2]int{x, int(queue.Probability(float64(x)) * 1000)}
		data = append(data, prob)
	}

	fmt.Println("expected length of system\n\t=> Ls =", Queue.Round(queue.LengthOfSystem()))
	fmt.Println("expected length of Queue\n\t=> Lq =", Queue.Round(queue.LengthOfQueue()))

	waitingTimeInQueue := queue.WaitingTimeInQueue()
	waitingTimeInSystem := queue.WaitingTimeInSystem()

	fmt.Println("expected waiting time in Queue = ", waitingTimeInQueue, "\n\t=> Wq =", Queue.Round(waitingTimeInQueue*60), "minutes")
	fmt.Println("expected waiting time in system = ", waitingTimeInSystem, "\n\t=> Ws =", Queue.Round(waitingTimeInSystem*60), "minutes \n ")

	plot := barchart.BarChartXYs(data)
	if err := barchart.Fprint(os.Stdout, plot, barchart.Linear(25)); err != nil {
		panic(err)
	}
}
