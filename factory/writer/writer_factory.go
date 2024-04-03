package writer

type customWriteError struct {
	errorMessage string
}

func GetWriter(writerConfig string, writerType string) (Writer, error) {
	if writerType == "gcs" {
		return GCSWriter{config: writerConfig}, nil
	} else {
		return nil, &customWriteError{errorMessage: "No Writer Implemenation Present"}
	}
}

func (c *customWriteError) Error() string {
	return c.errorMessage
}
