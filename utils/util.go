package utils

import (
	"encoding/binary"
	"errors"
	"fmt"
	"go-scheduler/types"
	"time"
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

func ComputeTime(d int, h int, m int) (time.Time, error) {
	var current = time.Now()
	if d < current.Day() {
		return current, errors.New("day given before present day")
	} else if d == current.Day() {
		if h < current.Hour() {
			return current, errors.New("hour given before present hour")
		}
	}

	timein := time.Date(current.Year(), current.Month(), d, h, m, 0, 0, time.UTC)
	return timein, nil
}
