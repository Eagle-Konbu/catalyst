package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Eagle-Konbu/catalyst/internal/usecase"
	"github.com/briandowns/spinner"
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
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Turning on the light..."
		s.Start()

		uc := getLightUsecaseOrExit(cmd)
		err := uc.TurnOnLight()

		s.Stop()
		fmt.Print("\r") // clear line

		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "failed to turn on light:", err)
			os.Exit(1)
		}
		fmt.Println("The light has been turned on successfully (´・ω・`)")
	},
}

var offCmd = &cobra.Command{
	Use:   "off",
	Short: "Turn the light off",
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		s.Suffix = " Turning off the light..."
		s.Start()

		uc := getLightUsecaseOrExit(cmd)
		err := uc.TurnOffLight()

		s.Stop()
		fmt.Print("\r") // clear line

		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "failed to turn off light:", err)
			os.Exit(1)
		}
		fmt.Println("The light has been turned off successfully (´・ω・`)")
	},
}

func init() {
	rootCmd.AddCommand(lightCmd)
	lightCmd.AddCommand(onCmd)
	lightCmd.AddCommand(offCmd)
}

func getLightUsecaseOrExit(cmd *cobra.Command) *usecase.LightUsecase {
	if lightId == "" || token == "" {
		fmt.Fprintln(cmd.ErrOrStderr(), "lightId or token is not set. Please check your config file or environment variables.")
		os.Exit(1)
	}
	return usecase.NewLightUsecase(lightId, token)
}
