package server

import (
	"github.com/spf13/cobra"
)

type cmdArgs struct {
	verbose bool
}

var rootCmd = NewRootCmd()

func NewRootCmd() *cobra.Command {

	var args cmdArgs

	cmd := &cobra.Command{
		Use:   "music-library-server",
		Short: "server of music-library",
	}

	cmd.PersistentFlags().BoolVar(&args.verbose, "verbose", false, "verbose mode")

	cmd.AddCommand(StartCmd)

	return cmd
}

func Execute() error {
	return rootCmd.Execute()
}
