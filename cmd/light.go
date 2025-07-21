/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Eagle-Konbu/catalyst/internal/usecase"
	"github.com/spf13/cobra"
)

// lightCmd represents the light command
var lightCmd = &cobra.Command{
	Use:   "light",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 || (args[0] != "on" && args[0] != "off") {
			return fmt.Errorf("usage: light [on|off]")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if id == "" || token == "" {
			fmt.Fprintln(cmd.ErrOrStderr(), "id or token is not set. Please check your config file or environment variables.")
			os.Exit(1)
		}

		uc := usecase.NewLightUsecase(id, token)
		var err error
		if args[0] == "on" {
			err = uc.TurnOnLight()
		} else {
			err = uc.TurnOffLight()
		}
		if err != nil {
			fmt.Fprintln(cmd.ErrOrStderr(), "failed to switch light:", err)
			os.Exit(1)
		}
		fmt.Printf("light has been turned %s successfully :)\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(lightCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lightCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lightCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
