package cmd

import (
	"github.com/Eagle-Konbu/catalyst/internal/misc"
	"github.com/Eagle-Konbu/catalyst/internal/usecase"
	"github.com/spf13/cobra"
)

// lightCmd represents the light command
var lightCmd = &cobra.Command{
	Use:   "light",
	Short: "Control the light device",
	Long:  `Control the light device using subcommands (on/off).`,
}

var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Turn the light on",
	Run: func(cmd *cobra.Command, args []string) {
		s := misc.NewSpinner(" Turning on the light...")
		uc := getLightUsecaseOrExit(cmd)
		err := uc.TurnOnLight()

		s.Stop()
		misc.PrintClearLine()

		if err != nil {
			misc.PrintErrorAndExit(cmd, "failed to turn on light", err)
		}
		misc.PrintSuccess("The light has been turned on successfully (´・ω・`)")
	},
}

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn the light off",
	Run: func(cmd *cobra.Command, args []string) {
		s := misc.NewSpinner(" Turning off the light...")
		uc := getLightUsecaseOrExit(cmd)
		err := uc.TurnOffLight()

		s.Stop()
		misc.PrintClearLine()

		if err != nil {
			misc.PrintErrorAndExit(cmd, "failed to turn off light", err)
		}
		misc.PrintSuccess("The light has been turned off successfully (´・ω・`)")
	},
}

func init() {
	rootCmd.AddCommand(lightCmd)
	lightCmd.AddCommand(onCmd)
	lightCmd.AddCommand(offCmd)
}

func getLightUsecaseOrExit(cmd *cobra.Command) *usecase.LightUsecase {
	if lightId == "" || token == "" {
		misc.PrintErrorAndExit(cmd, "lightId or token is not set. Please check your config file or environment variables.", nil)
	}
	return usecase.NewLightUsecase(lightId, token)
}
