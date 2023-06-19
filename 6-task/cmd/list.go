package cmd

import (
	"fmt"
	"os"

	"github.com/Gophercises/task/internals"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := internals.ListTasks()
		if err != nil {
			internals.Exitf("%v", err)
		}

		if len(tasks) == 0 {
			fmt.Println("there are no incompleted tasks")
			os.Exit(0)
		}

		fmt.Println("You have the following tasks:")
		for _, t := range tasks {
			fmt.Printf("%d. %s\n", t.ID, t.Title)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
