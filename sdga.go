/*
This file defines the core components of the State-Dependent Geometric Algebra (SDGA)
Simulation engine, based on the concepts from the PRISM project.

- Multivector: Represents a physical state (Ψ), defined by its Geometry and Energy.
- SDGA_Operator: An interface for dynamic transformations (like Genesis, Quench)
that operate on Multivectors over time.
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

// --- MULTIVECTOR DEFINTION AND METHODS ---

// Multivector represents a physical state (Ψ) in the SDGA framework.
// It's defined by its geometry (shape) and energy distribution.
type Multivector struct {
	Name string
	// Geometry is a 2D mask for the shape (alpha values 0-255)
	Geometry *image.Alpha
	// Energy is a 2D map of energy levels (grayscale v0-255)
	Energy *image.Gray
}

// newMultivector creates a new Multivector with zeroed (blank) images.
func newMultivector(name string, bounds image.Rectangle) Multivector {
	return Multivector{
		Name:     name,
		Geometry: image.NewAlpha(bounds),
		Energy:   image.NewGray(bounds),
	}
}

// TotalEnergy calculates the total energy contained in the state.
// It only sums energy where geometry is present (Alpha > 0).
func (mv Multivector) TotalEnergy() uint64 {
	var total uint64
	bounds := mv.Energy.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Only count energy where geometry exists
			if mv.Geometry.AlphaAt(x, y).A > 0 {
				total += uint64(mv.Energy.GrayAt(x, y).Y)
			}
		}
	}
	return total
}

// String() provide human-readable representation for logging.
func (mv Multivector) String() string {
	return fmt.Sprintf("<Multivector: %s | Total Energy: %d | Shape: %s>",
		mv.Name, mv.TotalEnergy(), mv.Geometry.Bounds().Size())
}

// --- SDGA OPERATOR INTERFACE ---

// SDGA_Operator defines a transformation between states over time.
type SDGA_Operator interface {
	// Apply launches a goroutine for calculating transformation.
	// Returns a read-only channel streaming intermediate states.
	Apply(initial, target Multivector, durationSteps int) <-chan Multivector
}

// --- OPERATOR IMPLEMENTATIONS ---

// GenesisOperator (𝐺): Transforms a state towards the Active Multivector (Ψ₁).
// Models a controlled, linear powerup (Inclining/Expanding).
type GenesisOperator struct{}

func (g GenesisOperator) Apply(initial, target Multivector, durationSteps int) <-chan Multivector {
	out := make(chan Multivector)
	bounds := initial.Geometry.Bounds()

	go func() {
		defer close(out)
		for i := 0; i <= durationSteps; i++ {
			t := float64(i) / float64(durationSteps) // Linear interpolation factor (0.0 to 1.0)

			// Create new intermediate image buffers
			interGeom := image.NewAlpha(bounds)
			interEnergy := image.NewGray(bounds)

			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					// Linear Interpolation (Lerp) Geometry
					geomA := float64(initial.Geometry.AlphaAt(x, y).A)
					geomB := float64(target.Geometry.AlphaAt(x, y).A)
					interGeom.SetAlpha(x, y, color.Alpha{A: uint8(geomA*(1-t) + geomB*t)})

					// Lerp for Energy
					energyA := float64(initial.Energy.GrayAt(x, y).Y)
					energyB := float64(target.Energy.GrayAt(x, y).Y)
					interEnergy.SetGray(x, y, color.Gray{Y: uint8(energyA*(1-t) + energyB*t)})
				}
			}
			out <- Multivector{
				Name:     fmt.Sprintf("Genesis-%.0f%%", t*100),
				Geometry: interGeom,
				Energy:   interEnergy,
			}
		}
	}()

	return out
}

// QuenchingOperator (𝑄): Transforms a state towards the Null Multivector (Ψ₀).
// Models a rapid, non-linear energy dissipation (Declining/Converging).
type QuenchingOperator struct{}

func (q QuenchingOperator) Apply(initial, target Multivector, durationSteps int) <-chan Multivector {
	out := make(chan Multivector)
	bounds := initial.Geometry.Bounds()

	go func() {
		defer close(out)

		for i := 0; i <= durationSteps; i++ {
			t := float64(i) / float64(durationSteps) // 0.0 to 1.0
			decayFactor := math.Pow(1.0-t, 3.0)      // Non-linear decay (t³)

			// Intermediate image buffers
			interGeom := image.NewAlpha(bounds)
			interEnergy := image.NewGray(bounds)

			for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					// Interpolate from target to initial using the decay factor
					geomA := float64(target.Geometry.AlphaAt(x, y).A)
					geomB := float64(initial.Geometry.AlphaAt(x, y).A)
					interGeom.SetAlpha(x, y, color.Alpha{A: unint8(geomA + (geomB-geomA)*decayFactor)})

					energyA := float64(target.Energy.GrayAt(x, y).Y)
					energyB := float64(initial.Energy.GrayAt(x, y).Y)
					interEnergy.SetGray(x, y, color.Gray{Y: unint8(energyA + (energyB-energyA)*decayFactor)})
				}
			}

			out <- Multivector{
				Name:     Sprintf("Quench-%.0f%%", t*100),
				Geometry: interGeom,
				Energy:   interEnergy,
			}
		}
	}()
	return out
}

// PotentialityOperator (𝑃): Instantly establishes the 'Standby' Multivector (Ψₚ).
type PotentialityOperator struct{}

func (p PotentialityOperator) Apply(initial, target Multivector, durationSteps int) <-chan Multivector {
	out := make(chan Multivector)
	go func() {
		defer close(out)
		// This is an instant change.
		// Copy created to avoid data races if original target is modified.
		targetCopy := newMultivector(target.Name, target.Geometry.Bounds())
		copy(targetCopy.Geometry.Pix, target.Geometry.Pix)
		copy(targetCopy.Energy.Pix, target.Energy.Pix)

		out <- targetCopy
	}()
	return out
}
