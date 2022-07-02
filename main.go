package main

import (
	"fmt"
	"os"
)

// Image

const (
	imageHeight = 256
	imageWidth = 256
)

// Render
func main() {

	fmt.Fprintf(os.Stdout, "P3\n%d %d\n255\n", imageWidth, imageHeight)
	
	for j := imageHeight -1; j >= 0; j-- {
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / (imageWidth - 1)
			g := float64(j) / (imageHeight - 1)
			b := 0.25

			ir := int(r * 255.999)
			ig := int(g * 255.999)
			ib := int(b * 255.999)

			fmt.Fprintf(os.Stdout, "%d %d %d\n", ir, ig, ib)
		}
	}
}
