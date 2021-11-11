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
	Short: "Provides only an update-css endpoint",
	Run:   runUpdateCss,
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
