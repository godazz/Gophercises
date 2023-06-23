package cmd

import (
	"fmt"
	"strconv"

	"github.com/godazz/Gophercises/task/internals"
	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internals.Exitf("Missing Task ID")
		}

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			internals.Exitf("%v", err)
		}
		taskTitle, err := internals.MarkTaskAsCompleted(taskID)
		if err != nil {
			internals.Exitf("%v", err)
		}
		fmt.Printf("You have completed the %q task.", taskTitle)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
