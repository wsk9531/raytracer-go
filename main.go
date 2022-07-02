package main

import (
	"fmt"
	"os"
)

// Image

const IMAGE_HEIGHT int = 256
const IMAGE_WIDTH int = 256

// Render
func main() {

	fmt.Fprintf(os.Stdout, "P3\n%d %d\n255\n", IMAGE_WIDTH, IMAGE_HEIGHT)

	for j := IMAGE_HEIGHT - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d ", j)
		for i := 0; i < IMAGE_WIDTH; i++ {

			colour := NewColour(
				float64(i)/(float64(IMAGE_WIDTH-1)),
				float64(j)/(float64(IMAGE_HEIGHT-1)),
				0.25,
			)
			
			WriteColour(colour)
		}
	}
}
