package cmd

import (
	"bufio"
	"fmt"
	"github.com/charmbracelet/log"
	"io"
	"os"
	"strings"

	"github.com/osean-man/pinner/internal/database"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new pinned command",
	Run: func(cmd *cobra.Command, args []string) {
		command := getCommand(os.Stdin, args)

		existingPins, err := database.GetPins(db)
		if err != nil {
			log.Errorf("Error fetching pins: %s", err)
			os.Exit(1)
		}
		for _, pin := range existingPins {
			if pin.Command == command {
				log.Error("Error: A command with the same text already exists.")
				os.Exit(1)
			}
		}

		err = database.AddPin(db, command)
		if err != nil {
			log.Errorf("Error adding pin: %s", err)
			os.Exit(1)
		}

		log.Info("Command pinned successfully!")
	},
}

func getCommand(stdin io.Reader, args []string) string {
	fmt.Print("Enter the command you want to pin: ")
	reader := bufio.NewReader(stdin)
	command, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Errorf("Error reading input: %s", err)
		os.Exit(1)
	}

	return strings.TrimSpace(command)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
