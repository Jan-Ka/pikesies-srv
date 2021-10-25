package cmd

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Jan-Ka/pikesies-srv/handlers"
)

// retrieveCssCmd represents the retrieveCss command
var retrieveCssCmd = &cobra.Command{
	Use: "retrieveCss",
	Aliases: []string{
		"retrieve-css",
	},
	Short: "Provides a retrieve-css endpoint",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		startRetrieveCssEndpoint()
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

func startRetrieveCssEndpoint() {
	log.Info().Msg("retrieveCss called")

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HealthHandler)
	r.HandleFunc("/retrieve-css", RetrieveCssHandler)
	http.Handle("/", r)
	port := fmt.Sprintf(":%s", viper.GetString("port"))

	log.Info().Msgf("Listening on %s", port)

	http.ListenAndServe(port, r)
}

func RetrieveCssHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "Thanks for testing retrieve-css!")
}
