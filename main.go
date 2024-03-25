/*
This code demonstrates the implementation of a simple generative art system using Markov chains and quinary logic.

Generative art involves creating artwork using algorithms, often resulting in unpredictable and unique outcomes.
Markov chains are mathematical models that describe a sequence of events where the probability of each event
depends only on the state attained in the previous event. Quinary logic, a system with five possible states,
is used here to introduce variability into the generated artwork.

In this implementation, a Markov chain is utilized to determine the transitions between different states of the quinary logic system,
representing various visual elements or properties in the artwork. Each state corresponds to a different artistic element,
such as color, shape, or texture. By iterating through the Markov chain, the program generates a sequence of states,
which are then interpreted to produce the final artwork.

Through the combination of Markov chains and quinary logic, this code exemplifies how algorithmic techniques can be applied to
create visually captivating and dynamic generative art. The resulting artwork reflects the inherent complexity and richness that emerge
from the interaction of simple rules and randomness.
*/
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
)

// QuinaryLogic represents the five states of our quinary logic system.
type QuinaryLogic int

const (
	On QuinaryLogic = iota
	Off
	OnWithinOff
	OffWithinOn
	Neutral
)

// MarkovChain represents a simple Markov chain graphic generator.
type MarkovChain struct {
	transitionMatrix [][]QuinaryLogic
}

// NewMarkovChain creates a new Markov chain with the given transition matrix.
func NewMarkovChain(transitionMatrix [][]QuinaryLogic) *MarkovChain {
	return &MarkovChain{
		transitionMatrix: transitionMatrix,
	}
}

// GenerateGraphic generates a graphic using the Markov chain.
func (mc *MarkovChain) GenerateGraphic(width, height int) *gg.Context {
	dc := gg.NewContext(width, height)
	rand.Seed(time.Now().UnixNano())
	currentState := rand.Intn(len(mc.transitionMatrix))

	for y := 0; y < height; y += 50 {
		for x := 0; x < width; x += 50 {
			mc.drawShape(dc, x, y, currentState)
			currentState = mc.getNextState(currentState)
		}
	}

	return dc
}

// drawShape draws a shape based on the current state.
func (mc *MarkovChain) drawShape(dc *gg.Context, x, y, currentState int) {
	switch {
	case mc.transitionMatrix[currentState][0] == On && mc.transitionMatrix[currentState][1] == On && mc.transitionMatrix[currentState][2] == OffWithinOn:
		dc.DrawRectangle(float64(x), float64(y), 50, 50)
		dc.SetRGB(1, 0, 0)
	case mc.transitionMatrix[currentState][0] == Off && mc.transitionMatrix[currentState][1] == Off && mc.transitionMatrix[currentState][2] == OnWithinOff:
		dc.DrawCircle(float64(x)+25, float64(y)+25, 25)
		dc.SetRGB(0, 0, 1)
	case mc.transitionMatrix[currentState][0] == OnWithinOff && mc.transitionMatrix[currentState][1] == OnWithinOff && mc.transitionMatrix[currentState][2] == Neutral:
		dc.DrawRectangle(float64(x), float64(y), 50, 50)
		dc.SetRGB(0, 1, 0)
	case mc.transitionMatrix[currentState][0] == OffWithinOn && mc.transitionMatrix[currentState][1] == OffWithinOn && mc.transitionMatrix[currentState][2] == On:
		dc.DrawCircle(float64(x)+25, float64(y)+25, 25)
		dc.SetRGB(1, 1, 0)
	case mc.transitionMatrix[currentState][0] == Neutral && mc.transitionMatrix[currentState][1] == Neutral && mc.transitionMatrix[currentState][2] == Neutral:
		dc.DrawRectangle(float64(x), float64(y), 50, 50)
		dc.SetRGB(0, 0, 0)
	}
	dc.Fill()
}

// getNextState selects the next state based on the current state and transition probabilities.
func (mc *MarkovChain) getNextState(currentState int) int {
	return rand.Intn(len(mc.transitionMatrix[currentState]))
}

func main() {
	// Define transition matrix
	transitionMatrix := [][]QuinaryLogic{
		{On, On, OffWithinOn},
		{On, On, OnWithinOff},
		{OffWithinOn, OnWithinOff, Neutral},
	}

	// Create Markov chain
	mc := NewMarkovChain(transitionMatrix)

	// Generate graphic
	width, height := 400, 400
	dc := mc.GenerateGraphic(width, height)

	// Save graphic to file
	if err := dc.SavePNG("output.png"); err != nil {
		fmt.Println("Error saving graphic:", err)
		return
	}
	fmt.Println("Graphic generated and saved to output.png")
}
