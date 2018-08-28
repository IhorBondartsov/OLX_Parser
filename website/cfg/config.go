package cfg

const (
	Route = "0.0.0.0"
	Port  = 8000
)

type MicroServices struct{
	Host string
	Port string
	Prefix string
}

var(
	// userms
	UserMS = MicroServices{
		Port:"8001",
		Host:"http://127.0.0.1",
		Prefix:"API",
	}

	ParserMS= MicroServices{
		Port:"8002",
		Host:"http://127.0.0.1",
		Prefix:"API",
	}
)