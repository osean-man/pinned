package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/osean-man/pinner/internal/database"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a pinned command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "Error: You must provide the ID of the command to remove")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Invalid ID format")
			os.Exit(1)
		}

		err = database.RemovePin(db, id)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error removing pin:", err)
			os.Exit(1)
		}

		fmt.Println("Command removed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
