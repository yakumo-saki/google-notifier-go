package cmd

import (
	"github.com/spf13/cobra"
)

// broadcastCmd represents the broadcast command
var sendCmd = &cobra.Command{
	Use:   "send <mDNSname> <message>",
	Short: "Send to Google devices",
	Long: `Send notification to single device.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		send(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// broadcastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// broadcastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func send(mdnsName, message string) {
	l.Info().Msgf("call send by MDNS=%s MSG=%s", mdnsName, message)
}
