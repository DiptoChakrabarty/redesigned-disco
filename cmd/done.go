package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-scheduler/database"
	"os"
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
	tasks, err := database.GetAllTasks()
	if err != nil {
		fmt.Println("Some Error Occured ", err.Error())
		os.Exit(1)
	}

	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			fmt.Println("InValid Task Number")
			continue
		}
		task := tasks[id-1]
		err := database.DeleteTasks(task.Key)
		if err != nil {
			fmt.Printf("Falied to complete task \"%d Error %s\n", id, err)
		} else {
			fmt.Printf("Complete task \"%d", id)
		}
	}
	fmt.Println(ids)
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
