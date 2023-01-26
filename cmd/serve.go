package cmd

import (
	"fmt"
	"strconv"

	fileserver "github.com/dooomit/file-server/pkg/FileServer"
	"github.com/spf13/cobra"
)

func Serve(groupId string) *cobra.Command {
	return &cobra.Command{
		Use:     "serve <root-path> [port]",
		Long:    "Run this command in order to start a api server",
		GroupID: groupId,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := serveApi(args)
			if err != nil {
				fmt.Println(err.Error())
			}
		},
	}

}

func serveApi(args []string) error {
	port := "3000"
	rootPath := args[0]
	if len(args) == 2 {
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("[port] has to be a number, you provided: %s", args[0])
		}
		port = args[0]
	}
	fileServer := fileserver.NewFileServer(rootPath)
	return fileServer.Run(":" + port)
}
