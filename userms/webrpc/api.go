package webrpc

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
