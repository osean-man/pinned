package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/osean-man/pinner/internal/database"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a pinned command",
	Run: func(cmd *cobra.Command, args []string) {
		pins, err := database.GetPins(db)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error fetching pins:", err)
			os.Exit(1)
		}

		if len(pins) == 0 {
			fmt.Println("You have no pinned commands to update.")
			return
		}

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
			Selected: "🖊️  {{ .Command | red | cyan }}",
		}

		prompt := promptui.Select{
			Label:     "Select a command to update",
			Items:     items,
			Templates: templates,
			Size:      10,
			Searcher:  searcher,
		}

		index, _, err := prompt.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error selecting command to update:", err)
			os.Exit(1)
		}

		selectedCommandID := pins[index].ID

		promptForNewCommand := promptui.Prompt{
			Label:     "Enter the updated command",
			Default:   pins[index].Command,
			AllowEdit: true,
		}

		newCommand, err := promptForNewCommand.Run()
		if err != nil {
			log.Errorf("Error reading new command: %s", err)
			os.Exit(1)
		}

		err = database.UpdatePin(db, selectedCommandID, newCommand)
		if err != nil {
			log.Errorf("Error updating pin: %s", err)
			os.Exit(1)
		}

		log.Info("Command updated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
