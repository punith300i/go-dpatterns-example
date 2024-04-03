package reader

type Spark struct {
	Conn string
}

type Reader interface {
	Read(s Spark)
}
