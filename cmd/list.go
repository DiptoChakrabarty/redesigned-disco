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
	fmt.Println("list called")
	//group, _ := cmd.Flags().GetString("group")
	tasks, err := database.GetAllTasks()
	if err != nil {
		fmt.Println("Some Error Occured ", err.Error())
		os.Exit(1)
	}
	utils.ListAllTasksOfNamespace(tasks)

}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringP("group", "g", "default", "The group which you would like to know")
}
