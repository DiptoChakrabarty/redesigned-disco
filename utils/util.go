package utils

import (
	"encoding/binary"
	"fmt"
	"go-scheduler/types"
)

func ConvertIntToBytes(n int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(n))
	return b
}

func ConvertBytesToInt(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func ListAllTasksOfNamespace(tasks []types.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks left to complete in default task list")
		return
	}
	fmt.Println("Your Current Tasks Are : ")
	for i, tk := range tasks {
		fmt.Printf("%d : %s\n", i+1, tk.Value)
	}
}
