package cmd

import (
	"github.com/mizuki1412/go-core-kit/tool-local/frontdao"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(frontDaoCmd)
}

var frontDaoCmd = &cobra.Command{
	Use:   "genDao",
	Short: "generate front dao",
	Run: func(cmd *cobra.Command, args []string) {
		frontdao.Gen("http://127.0.0.1:8080/weather-server-go")
	},
}
