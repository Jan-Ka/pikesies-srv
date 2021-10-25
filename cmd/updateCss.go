package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCssCmd represents the updateCss command
var updateCssCmd = &cobra.Command{
	Use: "updateCss",
	Aliases: []string{
		"update-css",
	},
	Short: "Provides an update-css endpoint",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runUpdateCss,
}

func init() {
	rootCmd.AddCommand(updateCssCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCssCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCssCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runUpdateCss(cmd *cobra.Command, args []string) {
	fmt.Println("updateCss called")
}
