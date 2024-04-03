package reader

import "fmt"

type GCSReader struct {
	config string
}

func (GCSReader) Read(s Spark) {
	fmt.Println("Reading data from GCS")
}
