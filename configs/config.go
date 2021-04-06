package configs

var (
	//App
	Port          = ":8080"
	SentryDsn     = Vault("sentry_dsn").SentryDsn
	SecretSignMac = Vault("sign_mac").SecretSignMac

	//MailGun
	MgDomain        = Vault("mailgun").MgDomain
	MgPrivateAPIKey = Vault("mailgun").MgPrivateApiKey

	//SendPulse
	GrantType    = "client_credentials"
	ClientID     = Vault("sendpulse").SpClientId
	ClientSecret = Vault("sendpulse").SpClientSecret

	//Mjml
	MjmlApplicationId = Vault("mjml").MjmlApplicationId
	MjmlSecretKey     = Vault("mjml").MjmlSecretKey
)
