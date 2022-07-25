package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yakumo-saki/google-notifier-go/src/mdnsclient"
)

// broadcastCmd represents the broadcast command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan to Google devices on your network",
	Long: `For debug purpose or setting purpose.
	Scan for Google Home (mini) / Nest etc...`,
	Run: func(cmd *cobra.Command, args []string) {
		scan()
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// broadcastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// broadcastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func scan() {
	mdnsclient.Scan()
}
