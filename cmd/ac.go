package cmd

import (
	"fmt"
	"regexp"

	"github.com/Eagle-Konbu/catalyst/internal/misc"
	"github.com/Eagle-Konbu/catalyst/internal/usecase"
	"github.com/spf13/cobra"
)

// acCmd represents the ac command
var acCmd = &cobra.Command{
	Use:   "ac",
	Short: "Control the air conditioner mode and temperature",
	Long:  `Set the air conditioner to a specific mode (cool, dry, or warm) and temperature (16.0-30.0C in 0.5C increments) using subcommands.`,
}

var coolCmd = &cobra.Command{
	Use:   "cool [temperature]",
	Short: "Set the air conditioner to cool mode",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runAcSubcommand(cmd, "cool", args)
	},
}

var dryCmd = &cobra.Command{
	Use:   "dry [temperature]",
	Short: "Set the air conditioner to dry mode",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runAcSubcommand(cmd, "dry", args)
	},
}

var warmCmd = &cobra.Command{
	Use:   "warm [temperature]",
	Short: "Set the air conditioner to warm mode",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runAcSubcommand(cmd, "warm", args)
	},
}

func runAcSubcommand(cmd *cobra.Command, mode string, args []string) {
	re := regexp.MustCompile(`^(1[6-9]|2[0-9])(\.0|\.5)?$|^30(\.0)?$`)
	if !re.MatchString(args[0]) {
		misc.PrintErrorAndExit(cmd, "temperature must be 16.0 to 30.0 in 0.5 increments", nil)
	}
	var temp float64
	_, err := fmt.Sscanf(args[0], "%f", &temp)
	if err != nil {
		misc.PrintErrorAndExit(cmd, "temperature must be a number", err)
	}
	if acId == "" || token == "" {
		misc.PrintErrorAndExit(cmd, "acId or token is not set. Please check your config file or environment variables.", nil)
	}

	s := misc.NewSpinner("Updating air conditioner settings...")
	uc := usecase.NewAirconUsecase(acId, token)
	err = uc.SwitchAirconSettings(mode, temp)

	s.Stop()
	misc.PrintClearLine()

	if err != nil {
		misc.PrintErrorAndExit(cmd, "Failed to switch air conditioner settings", err)
	}
	misc.PrintSuccess("Air conditioner settings has been updated successfully (´・ω・`)")
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the current status of the air conditioner",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if acId == "" || token == "" {
			misc.PrintErrorAndExit(cmd, "acId or token is not set. Please check your config file or environment variables.", nil)
		}

		s := misc.NewSpinner(" Fetching status...")
		uc := usecase.NewAirconUsecase(acId, token)
		status, err := uc.GetAirconStatus()

		s.Stop()
		misc.PrintClearLine()
		if err != nil {
			misc.PrintErrorAndExit(cmd, "Failed to get air conditioner status", err)
		}
		fmt.Printf("Mode: %s, Temperature: %.1f\n", status.Mode, status.Temperature)
	},
}

func init() {
	rootCmd.AddCommand(acCmd)
	acCmd.AddCommand(coolCmd)
	acCmd.AddCommand(dryCmd)
	acCmd.AddCommand(warmCmd)
	acCmd.AddCommand(statusCmd)
}
