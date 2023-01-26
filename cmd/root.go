package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "filer-server",
	Short: "This file server serves files from a given directory",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	allGroupId := "general"
	userGroupId := "user"
	allGroup := &cobra.Group{ID: allGroupId, Title: allGroupId}
	userGroup := &cobra.Group{ID: userGroupId, Title: userGroupId}
	rootCmd.AddGroup(allGroup, userGroup)
	rootCmd.AddCommand(Serve(allGroupId))
}
