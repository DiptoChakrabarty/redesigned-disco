package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to task list",
	Run:   AddCmdImplement,
}

func AddCmdImplement(cmd *cobra.Command, args []string) {
	taskGiven := strings.Join(args, " ")
	fmt.Printf("Added \"%s\" to the tasks list.\n", taskGiven)
}

func init() {
	RootCmd.AddCommand(addCmd)
}
