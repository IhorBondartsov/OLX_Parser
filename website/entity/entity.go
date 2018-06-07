package entity


type EchoReq struct {
	Name string
}

type EchoRes struct {
	Answer string
}

type LoginReq struct {
	Login    string
	Password string
}

type LoginResp struct {
	RefreshToken string
	AccessToken  string
}

type AcessTokenRequest struct {
	RefreshToken string
}

type AcessTokenResponse struct {
	AcessToken string
}

type MakeOrderReq struct {
	Token     string
	URL       string
	PageLimit int
	Mail      string
	DateTo    int64
	Frequency int
	UserID    int
}

type MakeOrderRes struct {
}

type ShowAllOderReq struct {
	Token  string
	UserID int
}
type ShowAllOderResp struct {
	Orders []Order
}

type GetAdvertisementByOrderReq struct {
	Token   string
	OrderID int
}

type GetAdvertisementByOrderResp struct {
	Advertisements []Advertisement
}

type Advertisement struct {
	ID      int
	URL     string
	Title   string
	Time    int64
	OrderID int
}

type Order struct {
	ID             int
	UserID         int
	URL            string
	PageLimit      int
	Mail           string
	ExpirationTime int64
	Frequency      int
}
