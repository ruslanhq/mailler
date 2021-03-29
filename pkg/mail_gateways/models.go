package mail_gateways

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ErrorCode   int    `json:"error_code"`
}

type BalanceDetailed struct {
	Balance *struct {
		Main float32 `json:"main,omitempty"`
	} `json:"balance,omitempty"`
	Email *struct {
		EmailsLeft int `json:"emails_left,omitempty"`
	} `json:"email,omitempty"`
}

type Query struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}