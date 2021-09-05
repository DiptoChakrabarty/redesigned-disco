package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-scheduler/database"
	"os"
)

var listall_Cmd = &cobra.Command{
	Use:   "all",
	Short: "Lists All of the total tasks",
	Run:   ListAllTasks,
}

func ListAllTasks(cmd *cobra.Command, args []string) {
	err := database.GetAllTasks()
	if err != nil {
		fmt.Println("Some Error Occured ", err.Error())
		os.Exit(1)
	}
}

func init() {
	listCmd.AddCommand(listall_Cmd)
}
