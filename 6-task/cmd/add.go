package cmd

import (
	"fmt"
	"strings"

	"github.com/Gophercises/task/internals"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			internals.Exitf("Missing Task Title\n")
		}

		task := internals.Task{Title: strings.Join(args, " ")}

		if err := internals.CreateTask(&task); err != nil {
			internals.Exitf("%v", err)
		}
		fmt.Printf("Added %q to your task list.\n", task.Title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
