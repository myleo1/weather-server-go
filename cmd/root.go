package cmd

import (
	corekit "github.com/mizuki1412/go-core-kit/init"
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"github.com/mizuki1412/go-core-kit/service/configkit"
	"github.com/mizuki1412/go-core-kit/service/logkit"
	"github.com/mizuki1412/go-core-kit/service/restkit"
	"github.com/spf13/cobra"
	"linkortech/framework/lkt-service-core/profile"
	"weather-server-go/mod/weather/action/weather"
)

func init() {
	initkit.DefFlags(rootCmd)
	profile.DefFlags(rootCmd)
}

var rootCmd = &cobra.Command{
	Use: "weather-server-go",
	Run: func(cmd *cobra.Command, args []string) {
		initkit.BindFlags(cmd)
		/// mod
		restkit.AddActions(weather.Init)
		/// service init

		//taskStart()
		_ = restkit.Run()
	},
}

/// 系统需要的定时任务
func taskStart() {
	if configkit.GetBoolD(corekit.ConfigKeyProfileDev) {
		return
	}
	//cronkit.Scheduler().Start()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logkit.Fatal(err.Error())
	}
}
