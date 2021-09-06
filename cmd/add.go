package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-scheduler/database"
	"os"
	"strings"
	"time"
)

var current = time.Now()

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to task list",
	Run:   AddCmdImplement,
}

var timeAddCmd = &cobra.Command{
	Use:   "time",
	Short: "Specify time limit for the task",
	Run:   AddCmdTimeImplement,
}

func AddCmdImplement(cmd *cobra.Command, args []string) {
	taskGiven := strings.Join(args, " ")
	group, _ := cmd.Flags().GetString("group")

	_, err := database.AddTask(taskGiven, []byte(group))
	if err != nil {
		fmt.Println("Some Error Occured ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Added \"%s\" to the tasks list.\n", taskGiven)
}

func AddCmdTimeImplement(cmd *cobra.Command, args []string) {
	taskGiven := strings.Join(args, " ")
	group, _ := cmd.Flags().GetString("group")
	hour, _ := cmd.Flags().GetInt("hour")
	minute, _ := cmd.Flags().GetInt("minute")
	day, _ := cmd.Flags().GetInt("day")

	fmt.Println(day, hour, minute, group)

	fmt.Printf("Added \"%s\" to the tasks list.\n", taskGiven)
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringP("group", "g", "default", "The group which you would like to know")
	timeAddCmd.PersistentFlags().IntP("hour", "r", int(current.Hour()), "Specify hours for time limit")
	timeAddCmd.PersistentFlags().IntP("minute", "m", int(current.Minute()), "Specify minutes for time limit")
	timeAddCmd.PersistentFlags().IntP("day", "d", int(current.Day()), "Specify day for time limit")
	addCmd.AddCommand(timeAddCmd)
}
