package writer

type Spark struct {
	Conn string
}

type Writer interface {
	Write(s Spark)
}
