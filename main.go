/*
This program is the main controller for the Quinary Logic SDGA simulation.

- Test Phase 1: Runs a text-based simulation to validate the
concurrent SDGA_Operator logic.
- Test Phase 2: Renders the simulation history to a graphical output (PNG)
and displays it in an X-Window.
*/
package main

import "flag"

func main() {
	// --- CLI ---
	outputFile := flag.String("o", "output.png", "Output PNG file name; defaults to `output.png`. Use `NONE` to disable saving.")
	useDisplay := flag.Bool("display", true, "Display the graphic in an X window.")
	flag.Int("width", "800", "Image width (simulation X-axis).")
	flag.Int("height", "600", "Image height (simulation Time-axis).")
	flag.Parse()
}
