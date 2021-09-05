package cmd

import (
	"fmt"
	"go-scheduler/database"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to task list",
	Run:   AddCmdImplement,
}

func AddCmdImplement(cmd *cobra.Command, args []string) {
	taskGiven := strings.Join(args, " ")
	// group, _ := cmd.Flags().GetString("group")

	_, err := database.AddTask(taskGiven)
	if err != nil {
		fmt.Println("Some Error Occured ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Added \"%s\" to the tasks list.\n", taskGiven)
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringP("group", "g", "default", "The group which you would like to know")
}
