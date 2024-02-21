package cmd

import (
	"fmt"
	"github.com/charmbracelet/log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/osean-man/pinner/internal/database"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all pinned commands",
	Run: func(cmd *cobra.Command, args []string) {
		pins, err := database.GetPins(db)
		if err != nil {
			log.Errorf("Error fetching pins: %s", err)
			os.Exit(1)
		}

		if len(pins) == 0 {
			fmt.Println("No pinned commands found.")
			return
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"ID", "Command"})

		t.SetStyle(table.StyleLight)

		for _, pin := range pins {
			t.AppendRow([]interface{}{pin.ID, pin.Command})
		}

		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
