package misc

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

func NewSpinner(suffix string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = suffix
	s.Start()
	return s
}

func PrintErrorAndExit(cmd *cobra.Command, msg string, err error) {
	if err != nil {
		fmt.Fprintf(cmd.ErrOrStderr(), "%s: %v\n", msg, err)
	} else {
		fmt.Fprintln(cmd.ErrOrStderr(), msg)
	}
	os.Exit(1)
}

func PrintSuccess(msg string) {
	fmt.Println(msg)
}

func PrintClearLine() {
	fmt.Print("\r")
}
