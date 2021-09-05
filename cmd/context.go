package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-scheduler/database"
)

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Lists All the tasks",
	Run:   ContextCmdImplement,
}

func ContextCmdImplement(cmd *cobra.Command, args []string) {
	for _, arg := range args {
		database.CreateBucket([]byte(arg))
		fmt.Printf("NameSpace %s added", arg)
	}
}

func init() {
	RootCmd.AddCommand(contextCmd)
}
