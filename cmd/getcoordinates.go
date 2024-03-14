package cmd

import (
	"fmt"
	"os"

	"github.com/martoc/ipstack-cli/core"
	"github.com/martoc/ipstack-cli/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	getCoordinates.Flags().StringP("access-key", "a", "", "Access key")
	getCoordinates.Flags().StringP("ip", "p", ".", "IP address")
	viper.BindEnv("access-key", "ACCESS_KEY") //nolint:errcheck
}

var getCoordinates = &cobra.Command{
	Use:   "get-coordinates",
	Short: "Get coordinates from an IP address",
	Long:  `Get coordinates from an IP address`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := cmd.Flags().GetString("ip")
		accessKey, _ := cmd.Flags().GetString("access-key")
		result, err := (&core.GetCoordinatesCommandBuilder{}).SetContext(cmd.Context()).SetIP(ip).SetAccessKey(accessKey).Build().Execute()
		if err != nil {
			logger.GetInstance().Error(err)
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, result)
	},
}
