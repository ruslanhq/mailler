package mail_gateways

type oauthTokenResponse struct {
	AccessToken string `json:"access_token"`
	ErrorCode   int    `json:"error_code"`
}

type BalanceDetailed struct {
	Balance *struct {
		Main string `json:"main,omitempty"`
	} `json:"balance,omitempty"`
	Email *struct {
		EmailsLeft int `json:"emails_left,omitempty"`
	} `json:"email,omitempty"`
}

type Query struct {
	Name         string `json:"name"`
	Mail         string `json:"mail"`
	TemplateName string `json:"template_name"`
	Mac          string `json:"MAC"`
}

type BalanceInfo struct {
	DateCheckBalance string
	Balance          int
}
