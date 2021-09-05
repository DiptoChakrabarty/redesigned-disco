package cmd

import (
	"fmt"
	"go-scheduler/database"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists All the tasks",
	Run:   ListCmdImplement,
}

func ListCmdImplement(cmd *cobra.Command, args []string) {
	fmt.Println("list called")
	tasks, err := database.GetAllTasks()
	if err != nil {
		fmt.Println("Some Error Occured ", err.Error())
		os.Exit(1)
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks left to complete in default task list")
		return
	}
	fmt.Println("Your Current Tasks Are : ")
	for i, tk := range tasks {
		fmt.Printf("%d : %s\n", i+1, tk.Value)
	}
}

func init() {
	RootCmd.AddCommand(listCmd)
}
