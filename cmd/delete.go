package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/osean-man/pinner/internal/database"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a pinned command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Error: You must provide the ID of the command to delete")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Invalid ID format")
			os.Exit(1)
		}

		err = database.DeletePin(db, id) // Assuming you have DeletePin implemented
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error deleting pin:", err)
			os.Exit(1)
		}

		fmt.Println("Command deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
