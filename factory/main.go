package main

import (
	"factory/reader"
	"factory/writer"
	"fmt"
)

func main() {
	gbqReader, err := reader.GetReader("gbq conn string", "gbq")
	if err == nil {
		fmt.Println(gbqReader)
		gbqReader.Read(reader.Spark{Conn: "test"})
	}

	gcsReader, err := reader.GetReader("gcs conn string", "gcs")
	if err == nil {
		fmt.Println(gcsReader)
		gcsReader.Read(reader.Spark{Conn: "test"})
	}

	gcsWriter, err := writer.GetWriter("writer config", "gcs")
	if err == nil {
		fmt.Println(gcsWriter)
	}

}
