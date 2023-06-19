package cmd

import (
	"fmt"
	"os"

	"github.com/Gophercises/task/internals"
	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all of your complete tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := internals.ListTasks(true)
		if err != nil {
			internals.Exitf("%v", err)
		}

		if len(tasks) == 0 {
			fmt.Println("there are no completed tasks")
			os.Exit(0)
		}

		fmt.Println("You have finished the following tasks:")
		for _, t := range tasks {
			fmt.Printf("%d. %s\n", t.ID, t.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)
}
