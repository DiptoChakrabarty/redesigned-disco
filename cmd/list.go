package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-scheduler/database"
	"go-scheduler/utils"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists All the tasks",
	Run:   ListCmdImplement,
}

func ListCmdImplement(cmd *cobra.Command, args []string) {
	group, _ := cmd.Flags().GetString("group")
	all, _ := cmd.Flags().GetString("all")
	if all == "yes" {
		err := database.GetAllTasks()
		if err != nil {
			fmt.Println("Some Error Occured ", err.Error())
			os.Exit(1)
		}
		return
	}
	tasks, err := database.GetAllTasksOfGroup([]byte(group))
	if err != nil {
		fmt.Println("Some Error Occured ", err.Error())
		os.Exit(1)
	}
	utils.ListAllTasksOfNamespace(tasks)

}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringP("group", "g", "default", "The group which you would like to know")
	listCmd.PersistentFlags().StringP("all", "a", "no", "To check all tasks of groups options yes no")
}
