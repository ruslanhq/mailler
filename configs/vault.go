package configs

import (
	"github.com/getsentry/sentry-go"
	"github.com/riotgames/vault-go-client"
	"log"
	"os"
)

type VaultStruct struct {
	//Mailgun
	MgDomain        string `json:"domain"`
	MgPrivateApiKey string `json:"private_api_key"`

	//Sendpulse
	SpClientId     string `json:"client_id"`
	SpClientSecret string `json:"client_secret"`

	//Mjml
	MjmlApplicationId string `json:"application_id"`
	MjmlSecretKey     string `json:"secret_key"`

	SentryDsn     string `json:"dsn"`
	SecretSignMac string `json:"secret_sign_mac"`
}

func Vault(path string) *VaultStruct {
	appRoleID := os.Getenv("VAULT_APPROLE_ID")
	appRoleSecretID := os.Getenv("VAULT_APPROLE_SECRET_ID")
	appRoleAuthPath := os.Getenv("VAULT_APPROLE_AUTH_PATH")
	secretMountPath := os.Getenv("VAULT_SECRET_MOUNT_PATH")
	secretPath := path

	client, err := vault.NewClient(vault.DefaultConfig())
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal(err.Error())
	}

	if _, err = client.Auth.AppRole.Login(vault.AppRoleLoginOptions{
		RoleID:    appRoleID,
		SecretID:  appRoleSecretID,
		MountPath: appRoleAuthPath}); err != nil {
		sentry.CaptureException(err)
		log.Fatal(err.Error())
	}
	secrets := &VaultStruct{}

	if _, err = client.KV2.Get(vault.KV2GetOptions{
		MountPath:     secretMountPath,
		SecretPath:    secretPath,
		UnmarshalInto: secrets,
	}); err != nil {
		sentry.CaptureException(err)
		log.Fatal(err.Error())
	}

	return secrets
}
