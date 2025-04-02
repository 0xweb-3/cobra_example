package cmd

import (
	"github.com/0xweb-3/cobra_example/user"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User service commands",
}

var lunchCmd = &cobra.Command{
	Use:   "lunch",
	Short: "Start the user service",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		service := user.NewUserService()
		service.Start(ctx)
	},
}

func init() {
	userCmd.AddCommand(lunchCmd)
}
