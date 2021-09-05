package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists All the tasks",
	Run:   ListCmdImplement,
}

func ListCmdImplement(cmd *cobra.Command, args []string) {
	fmt.Println("list called")
}

func init() {
	RootCmd.AddCommand(listCmd)
}
