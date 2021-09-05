package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-scheduler/database"
	"os"
	"strings"
)

var r string
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists All the tasks",
	Run:   ListCmdImplement,
}

var checkCmd = &cobra.Command{
	Use: "check",
	Run: Check,
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

func Check(cmd *cobra.Command, args []string) {
	b, _ := cmd.Flags().GetString("block")
	fmt.Println("The check command called")
	taskGiven := strings.Join(args, " ")
	fmt.Println(taskGiven, b)
}

func init() {
	RootCmd.AddCommand(listCmd)
	RootCmd.AddCommand(checkCmd)
	checkCmd.PersistentFlags().StringP("block", "b", "normal", "Just checking default")
}
