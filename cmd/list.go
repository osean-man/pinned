package cmd

import (
	"fmt"
	"os"

	"github.com/osean-man/pinned/internal/database"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all pinned commands",
	Run: func(cmd *cobra.Command, args []string) {
		pins, err := database.GetPins(db)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error fetching pins:", err)
			os.Exit(1)
		}

		if len(pins) == 0 {
			fmt.Println("No pinned commands found.")
			return
		}

		for _, pin := range pins {
			fmt.Printf("ID: %d, Command: %s\n", pin.ID, pin.Command)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
