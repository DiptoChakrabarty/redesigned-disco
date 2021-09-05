package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listall_Cmd = &cobra.Command{
	Use:   "all",
	Short: "Lists All of the total tasks",
	Run:   ListAllTasks,
}

func ListAllTasks(cmd *cobra.Command, args []string) {
	fmt.Println("list all commands called")
}

func init() {
	listCmd.AddCommand(listall_Cmd)
}
