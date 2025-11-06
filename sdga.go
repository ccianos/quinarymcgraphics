/*
This file defines the core components of the State-Dependent Geometric Algebra (SDGA)
Simulation engine, based on the concepts from the PRISM project.

- Multivector: Represents a physical state (Ψ), defined by its Geometry and Energy.
- SDGA_Operator: An interface for dynamic transformations (like Genesis, Quench)
that operate on Multivectors over time.
*/
package main

import "image"

// Multivector represents a physical state (Ψ) in the SDGA framework.
// It's defined by its geometry (shape) and energy distribution.
type Multivector struct {
	Name string
	// Geometry is a 2D mask for the shape (alpha values 0-255)
	Geometry *image.Alpha
	// Energy is a 2D map of energy levels (grayscale v0-255)
	Energy *image.Gray
}
