package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/osean-man/pinned/internal/database"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a pinned command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, "Error: You must provide the ID of the command to update and the new command")
			os.Exit(1)
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Invalid ID format")
			os.Exit(1)
		}

		newCommand := args[1]

		err = database.UpdatePin(db, id, newCommand)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error updating pin:", err)
			os.Exit(1)
		}

		fmt.Println("Command updated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
