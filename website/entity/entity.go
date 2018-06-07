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