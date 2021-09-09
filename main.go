package main

import (
	"github.com/DiptoChakrabart/task-manager/cmd"
	"github.com/DiptoChakrabart/task-manager/database"
	"log"
	"os"
	"path/filepath"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	DbPath := filepath.Join(homeDir, "tasks.db")
	err = database.Init(DbPath)
	if err != nil {
		log.Fatal(err)
	}
	cmd.RootCmd.Execute()

}
