package writer

import "fmt"

type GCSWriter struct {
	config string
}

func (GCSWriter) Write(s Spark) {
	fmt.Println("Writing data to GCS")
}
