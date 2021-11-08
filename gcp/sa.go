package gcp

import (
	"context"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

var lock = &sync.Mutex{}

type secretManager struct {
	waAppKey string
}

var secretManagerInstance *secretManager

func GetSecretManager() *secretManager {
	if secretManagerInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		secretManagerInstance = &secretManager{}
	}

	return secretManagerInstance
}

func (sa *secretManager) GetWaAppKey() (string, error) {
	if len(sa.waAppKey) <= 0 {
		lock.Lock()
		defer lock.Unlock()

		fnContext := context.Background()

		fnLog := log.With().Str("package", "gcp").Logger()

		saConfigPath := viper.GetString("gcp_service_account_path")
		saPath := saConfigPath

		if !filepath.IsAbs(saConfigPath) {
			basePath, _ := os.Getwd()

			saPath = path.Join(basePath, saConfigPath)
		}

		fnLog.Debug().Msgf("Reading GCP service account from %s", saPath)

		client, err := secretmanager.NewClient(fnContext, option.WithCredentialsFile(saPath))
		if err != nil {
			return "", fmt.Errorf("failed to create secret manager due to %s", err)
		}

		req := &secretmanagerpb.AccessSecretVersionRequest{
			Name: viper.GetString("wa_app_key_secret_key"),
		}

		result, err := client.AccessSecretVersion(fnContext, req)
		if err != nil {
			return "", fmt.Errorf("failed to access secret version due to %s", err)

		}

		sa.waAppKey = string(result.Payload.Data)
	}

	return sa.waAppKey, nil
}
