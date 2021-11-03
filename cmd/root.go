package cmd

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pikesies-srv",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdCtx := cmd.Context()
		ctxWaitGroup := cmdCtx.Value(CtxWaitGroupKey{}).(*sync.WaitGroup)

		ctxWaitGroup.Done()
		// send sigint to stop app
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ctx context.Context) {
	cobra.CheckErr(rootCmd.ExecuteContext(ctx))
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ${cwd}/.pikesies-srv.yaml)")

	rootCmd.PersistentFlags().String("port", "", "port to listen on (default is 8080)")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("port", "8080")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// run path
		exPath, err := os.Executable()
		cobra.CheckErr(err)
		exePath := filepath.Dir(exPath)

		viper.AddConfigPath(exePath)

		// home path
		homePath, err := os.UserHomeDir()

		cobra.CheckErr(err)

		viper.AddConfigPath(homePath)

		viper.SetConfigType("yaml")
		viper.SetConfigName(".pikesies-srv")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Info().Msgf("Using config file: %v", viper.ConfigFileUsed())
	} else {
		log.Info().Msg("No config file loaded")
	}
}
