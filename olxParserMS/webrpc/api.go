package webrpc

type EchoReq struct {
	Name string
}

type EchoRes struct {
	Answer string
}

type MakeOrderReq struct {
	Token          string
	URL            string
	PageLimit      int
	DeliveryMethod string
	DateTo         int64
	Frequency      int
	UserID         int
}

type MakeOrderRes struct {
}
