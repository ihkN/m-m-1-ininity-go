package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Queue struct {
	vec []float64
}

func (q *Queue) Dequeue() float64 {
	i := len(q.vec) - 1
	elem := q.vec[i]
	q.vec = q.vec[:i]
	return elem
}

func (q *Queue) Enqueue(x float64) {
	q.vec = append(q.vec, x)
}

func (q *Queue) Length() int {
	return len(q.vec)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	lambda := 8.0
	mu := 9.0 / 0.6

	nextArrival := getExponentialRandNum(lambda)
	nextDeparture := math.Inf(1)

	totalWait := 0.0
	served := 0

	queue := Queue{[]float64{}}

	for {
		//fmt.Println("nextArrival", nextArrival, "\nnextDeparture", nextDeparture)
		if nextArrival <= nextDeparture {
			if queue.Length() == 0 {
				nextDeparture = nextArrival + getExponentialRandNum(mu)
			}
			fmt.Println("Adding to Queue", nextArrival)
			queue.Enqueue(nextArrival)
			nextArrival += getExponentialRandNum(lambda)
		} else {
			elem := queue.Dequeue()
			fmt.Println("Removed from Queue", elem)
			wait := nextDeparture - elem
			fmt.Println("wait:", wait)
			totalWait += wait
			served++

			if served == 1000 {
				fmt.Println("Done!")
				fmt.Println("Total wait:", totalWait)

				return
			}
			if queue.Length() == 0 {
				nextDeparture = math.Inf(1)
			} else {
				nextDeparture += getExponentialRandNum(mu)
			}
		}
	}
}

func getExponentialRandNum(lambda float64) float64 {
	return math.Log(1-rand.Float64()) / (-lambda)
}

// -----------------------------------------------------------------------------

type Job struct {
}
