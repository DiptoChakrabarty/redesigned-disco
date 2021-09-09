package cmd

import (
	"fmt"
	"github.com/DiptoChakrabart/task-manager/database"
	"github.com/DiptoChakrabart/task-manager/utils"
	"github.com/spf13/cobra"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists All the tasks",
	Run:   ListCmdImplement,
}
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

func ListCmdImplement(cmd *cobra.Command, args []string) {
	group, _ := cmd.Flags().GetString("group")
	//all, _ := cmd.Flags().GetString("all")
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
	listCmd.AddCommand(listall_Cmd)
}
