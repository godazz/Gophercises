package cmd

import (
	"fmt"
	"strconv"

	"github.com/godazz/Gophercises/task/internals"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "delte a task from your TODO list",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internals.Exitf("Missing Task ID")
		}

		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			internals.Exitf("%v", err)
		}
		taskTitle, err := internals.DeleteTask(taskID)
		if err != nil {
			internals.Exitf("%v", err)
		}
		fmt.Printf("You have deleted the %q task.", taskTitle)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
