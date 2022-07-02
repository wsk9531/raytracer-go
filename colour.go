package main

import (
	"fmt"
	"os"
)

type colour vec3

func NewColour(r, g, b float64) colour {
	return colour(newVec3(r, g, b))
}

func WriteColour(c colour) {
	ir := int(c.X * 255.999)
	ig := int(c.Y * 255.999)
	ib := int(c.Z * 255.999)

	fmt.Fprintf(os.Stdout, "%d %d %d\n", ir, ig, ib)
}
