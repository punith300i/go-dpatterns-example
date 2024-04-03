package reader

type customError struct {
	s string
}

func GetReader(readerConfig, readerType string) (Reader, error) {
	if readerType == "gbq" {
		return GBQReader{config: readerConfig}, nil
	} else if readerType == "gcs" {
		return GCSReader{config: readerConfig}, nil
	} else {
		return nil, &customError{s: "No Factory implementation present"}
	}
}

func (c *customError) Error() string {
	return c.s
}
