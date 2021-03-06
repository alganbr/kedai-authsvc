package cmd

import (
	"github.com/alganbr/kedai-authsvc/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var RootCmd = &cobra.Command{
	Use:   "authsvc",
	Short: "User service API",
	Long:  "User service API",
	RunE: func(cmd *cobra.Command, args []string) error {
		fx.New(server.Module).Run()
		return nil
	},
}
