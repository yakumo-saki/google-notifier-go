package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yakumo-saki/google-notifier-go/src/mdnsclient"
)

// broadcastCmd represents the broadcast command
var broadcastCmd = &cobra.Command{
	Use:   "broadcast <messages>",
	Short: "Broadcast to all Google devices",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		broadcast()
	},
}

func init() {
	rootCmd.AddCommand(broadcastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// broadcastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// broadcastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func broadcast() {
	mdnsclient.Scan(false)
}
