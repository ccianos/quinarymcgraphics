/*
This program is the main controller for the Quinary Logic SDGA simulation.

- Test Phase 1: Runs a text-based simulation to validate the
concurrent SDGA_Operator logic.
- Test Phase 2: Renders the simulation history to a graphical output (PNG)
and displays it in an X-Window.
*/
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
)

func main() {
	// --- CLI ---
	outputFile := flag.String("o", "output.png", "Output PNG file name; defaults to `output.png`. Use `NONE` to disable saving.")
	useDisplay := flag.Bool("display", true, "Display the graphic in an X window.")
	flag.Int("width", "800", "Image width (simulation X-axis).")
	flag.Int("height", "600", "Image height (simulation Time-axis).")
	flag.Parse()

	fmt.Println("--- Phase 1 (text simulation): Initializing SDGA System ---")

	// --- Define Stable States ---

	// 2D spatial dimension of simulation
	simShape := image.Rect(0, 0, 100, 10) // width: 100, height: 10

	// PSI_NULL (Ψ₀): The Inactive state.
	// newMultivector Ψ₀ initialized with 0 energy and 0 geometry (transparent).
	PSI_NULL := newMultivector("Null (Ψ₀)", simShape)

	// PSI_ACTIVE (Ψ₁): The Active state.
	// newMultivector Ψ₁ initialized with full geometry and high energy.
	PSI_ACTIVE := newMultivector("Active (Ψ₁)", simShape)
	activeGeomY := simShape.Dy() / 2 // Center 2-pixel element
	for x := 0; x < simShape.Dx(); x++ {
		// Set Geometry (Alpha)
		PSI_ACTIVE.Geometry.SetAlpha(x, activeGeomY-1, color.Alpha{A: 255})
		PSI_ACTIVE.Geometry.SetAlpha(x, activeGeomY, color.Alpha{A: 255})
		// Set Energy (Grayscale)
		PSI_ACTIVE.Energy.SetGray(x, activeGeomY-1, color.Gray{Y: 250}) // 250 out of 255
		PSI_ACTIVE.Energy.SetGray(x, activeGeomY, color.Gray{Y: 250})
	}

	// PSI_POTENTIAL (Ψₚ): The Potential state.
}
