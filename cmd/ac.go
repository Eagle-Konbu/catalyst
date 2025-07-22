/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"

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
		fmt.Fprintln(cmd.ErrOrStderr(), "temperature must be 16.0 to 30.0 in 0.5 increments")
		os.Exit(1)
	}
	var temp float64
	_, err := fmt.Sscanf(args[0], "%f", &temp)
	if err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), "temperature must be a number")
		os.Exit(1)
	}
	if acId == "" || token == "" {
		fmt.Fprintln(cmd.ErrOrStderr(), "acId or token is not set. Please check your config file or environment variables.")
		os.Exit(1)
	}
	uc := usecase.NewAirconUsecase(acId, token)
	if err := uc.SwitchAirconSettings(mode, temp); err != nil {
		fmt.Fprintln(cmd.ErrOrStderr(), "Failed to switch air conditioner settings:", err)
		os.Exit(1)
	}
	fmt.Println("Air conditioner settings has been updated successfully!")
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the current status of the air conditioner",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if acId == "" || token == "" {
			fmt.Fprintln(cmd.ErrOrStderr(), "acId or token is not set. Please check your config file or environment variables.")
			os.Exit(1)
		}
		uc := usecase.NewAirconUsecase(acId, token)
		status, err := uc.GetAirconStatus()
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "Failed to get air conditioner status:", err)
			os.Exit(1)
		}

		var tempUnit string
		switch status.TempUnit {
		case "c":
			tempUnit = "℃"
		case "f":
			tempUnit = "°F"
		default:
			tempUnit = ""
		}

		fmt.Printf("Mode: %s, Temperature: %.1f%s\n", status.Mode, status.Temperature, tempUnit)
	},
}

func init() {
	rootCmd.AddCommand(acCmd)
	acCmd.AddCommand(coolCmd)
	acCmd.AddCommand(dryCmd)
	acCmd.AddCommand(warmCmd)
	acCmd.AddCommand(statusCmd)
}
