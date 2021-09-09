package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-scheduler/database"
)

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Work with categories for the groups",
	Run:   ContextCmdImplement,
}

var addBucketCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a category for the groups",
	Run:   AddContext,
}

var DeleteBucketCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a category for the groups",
	Run:   DeleteContext,
}

func ContextCmdImplement(cmd *cobra.Command, args []string) {
	for _, arg := range args {
		database.CreateBucket([]byte(arg))
		fmt.Printf("NameSpace %s added\n", arg)
	}
}

func AddContext(cmd *cobra.Command, args []string) {
	for _, arg := range args {
		database.CreateBucket([]byte(arg))
		fmt.Printf("NameSpace %s added\n", arg)
	}
}

func DeleteContext(cmd *cobra.Command, args []string) {
	for _, arg := range args {
		database.DeleteABucket([]byte(arg))
		fmt.Printf("NameSpace %s removed\n", arg)
	}
}

func init() {
	RootCmd.AddCommand(contextCmd)
	addBucketCmd.PersistentFlags().StringP("add", "a", "", "The category which you would like to add")
	DeleteBucketCmd.PersistentFlags().StringP("delete", "d", "", "The category which you would like to remove")
	contextCmd.AddCommand(addBucketCmd)
	contextCmd.AddCommand(DeleteBucketCmd)
}
