package main

import (
	"github.com/cosmos/iavl/v2/cmd/gen"
	"github.com/spf13/cobra"
)

func RootCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "iavl",
		Short: "benchmark cosmos/iavl",
	}
	cmd.AddCommand(gen.Command())
	return cmd, nil
}