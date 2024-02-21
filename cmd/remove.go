package cmd

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"strings"

	"github.com/osean-man/pinner/internal/database"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a pinned command",
	Run: func(cmd *cobra.Command, args []string) {
		pins, err := database.GetPins(db)
		if err != nil {
			fmt.Println("Error fetching pins:", err, os.Stderr)
			os.Exit(1)
		}

		if len(pins) == 0 {
			fmt.Println("You have no pinned commands to remove.")
			return
		}

		// Display Menu
		items := make([]Pin, len(pins))
		for i, p := range pins {
			items[i] = Pin{ID: p.ID, Command: p.Command}
		}

		searcher := func(input string, index int) bool {
			pin := items[index]
			return strings.Contains(strings.ToLower(pin.Command), strings.ToLower(input))
		}

		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "📌 {{ .Command | green }}",
			Inactive: "  {{ .Command }}",
			Selected: "✖️ {{ .Command | red | cyan }}",
		}

		prompt := promptui.Select{
			Label:     "Select a command to remove",
			Items:     items,
			Templates: templates,
			Size:      10,
			Searcher:  searcher,
		}

		index, _, err := prompt.Run()
		if err != nil {
			fmt.Println("Error selecting command:", err, os.Stderr)
			os.Exit(1)
		}

		selectedCommandID := pins[index].ID

		// Confirm Removal
		fmt.Printf("Are you sure you want to remove command '%s'? (y/N): ", pins[index].Command)

		reader := bufio.NewReader(os.Stdin)
		confirmation, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading confirmation:", err, os.Stderr)
			os.Exit(1)
		}

		confirmation = strings.TrimSpace(confirmation)

		if strings.ToLower(confirmation) != "y" {
			fmt.Println("Removal cancelled.")
			return
		}

		// Proceed with Removal
		err = database.RemovePin(db, selectedCommandID)
		if err != nil {
			fmt.Println("Error removing pin:", err, os.Stderr)
			os.Exit(1)
		}

		fmt.Println("Command removed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
