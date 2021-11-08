package cmd

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Jan-Ka/pikesies-srv/gcp"
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
		runLog := log.With().Str("package", "cmd").Str("cmd", "retrieveCss").Logger()

		port := fmt.Sprintf(":%s", viper.GetString("port"))

		cmdCtx := cmd.Context()
		ctxWaitGroup := cmdCtx.Value(CtxWaitGroupKey{}).(*sync.WaitGroup)

		sa := gcp.GetSecretManager()

		waAppKey, err := sa.GetWaAppKey()

		if err != nil {
			runLog.Error().Msg(err.Error())
			return
		}

		runLog.Debug().Msgf("Got Secret with length of %v\n", len(waAppKey))

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
