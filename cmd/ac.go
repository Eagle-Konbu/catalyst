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
	Long: `Set the air conditioner to a specific mode (cool, dry, or warm) and temperature (16.0-30.0C in 0.5C increments).

Example:
  catalyst ac cool 24.5
This sets the air conditioner to cool mode at 24.5C.`,
	Args: cobra.MatchAll(cobra.ExactArgs(2), func(cmd *cobra.Command, args []string) error {
		mode := args[0]
		switch mode {
		case "cool", "dry", "warm":
			// valid modes
		default:
			return fmt.Errorf("mode must be one of 'cool', 'dry', or 'warm'")
		}
		re := regexp.MustCompile(`^(1[6-9]|2[0-9])(\.0|\.5)?$|^30(\.0)?$`)
		if !re.MatchString(args[1]) {
			return fmt.Errorf("temperature must be 16.0 to 30.0 in 0.5 increments")
		}
		var temp float64
		_, err := fmt.Sscanf(args[1], "%f", &temp)
		if err != nil {
			return fmt.Errorf("temperature must be a number")
		}
		return nil
	}),
	Run: func(cmd *cobra.Command, args []string) {
		mode := args[0]
		var temp float64
		fmt.Sscanf(args[1], "%f", &temp)

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
	},
}

func init() {
	rootCmd.AddCommand(acCmd)
}
