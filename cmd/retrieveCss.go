package cmd

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"sync"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"

	"github.com/Jan-Ka/pikesies-srv/server"
)

// retrieveCssCmd represents the retrieveCss command
var retrieveCssCmd = &cobra.Command{
	Use: "retrieveCss",
	Aliases: []string{
		"retrieve-css",
	},
	Short: "Provides a retrieve-css endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		runLog := log.With().Str("cmd", "retrieveCss").Logger()

		port := fmt.Sprintf(":%s", viper.GetString("port"))

		cmdCtx := cmd.Context()
		ctxWaitGroup := cmdCtx.Value(CtxWaitGroupKey{}).(*sync.WaitGroup)

		saConfigPath := viper.GetString("gcp_service_account_path")
		saPath := saConfigPath

		if !filepath.IsAbs(saConfigPath) {
			basePath, _ := os.Getwd()

			saPath = path.Join(basePath, saConfigPath)
		}

		runLog.Debug().Msgf("Reading GCP service account from %s", saPath)

		client, err := secretmanager.NewClient(cmdCtx, option.WithCredentialsFile(saPath))
		if err != nil {
			runLog.Error().Msgf("Failed to create secret manager due to %s\n", err)
			return
		}

		req := &secretmanagerpb.AccessSecretVersionRequest{
			Name: viper.GetString("wa_app_key_secret_key"),
		}

		result, err := client.AccessSecretVersion(cmdCtx, req)
		if err != nil {
			runLog.Error().Msgf("Failed to access secret version due to %s\n", err)
			return
		}

		runLog.Debug().Msgf("Got Secret with length of %v\n", len(string(result.Payload.Data)))

		go server.RunServer(cmdCtx, ctxWaitGroup, port, "/retrieve-css", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "Thanks for testing retrieve-css!")
		})
	},
}

func init() {
	rootCmd.AddCommand(retrieveCssCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// retrieveCssCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// retrieveCssCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
