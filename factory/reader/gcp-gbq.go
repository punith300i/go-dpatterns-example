package reader

import "fmt"

type GBQReader struct {
	config string
}

func (GBQReader) Read(s Spark) {
	fmt.Println("Reading data from GBQ")
}
