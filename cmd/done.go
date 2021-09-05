package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Completes a task in task list",
	Run:   DoneCmdImplement,
}

func DoneCmdImplement(cmd *cobra.Command, args []string) {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("Failed to parse given argument", arg)
		} else {
			ids = append(ids, id)
		}
	}
	fmt.Println(ids)
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
