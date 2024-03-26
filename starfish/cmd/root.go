package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:        "starfish",
	Short:      "starfish app",
	Long:       `The starfish program implements scanning of contracts and main chain currency transfer information on the eth chain through wss and rpc.`,
	SuggestFor: []string{},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("daemon", false, "Run in daemon mode")
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
