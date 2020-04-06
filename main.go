package main

import (
	"fmt"
	"github.com/paulsmith/gogeos/geos"
)

func main() {
	var (
		a = geos.Must(geos.FromWKT("POLYGON ((0 3, 2 3, 3 1, 1 0, 0 1.5, 0 3))"))
		b = geos.Must(geos.FromWKT("POLYGON ((1 2, 1.5 4, 3.5 4, 4.5 2.5, 1 2))"))
	)
	c := geos.Must(a.Intersection(b))
	fmt.Println(c)
}
