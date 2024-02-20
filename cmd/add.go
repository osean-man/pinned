package cmd

import (
	"bufio"
	"fmt"
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
			fmt.Fprintln(os.Stderr, "Error fetching pins:", err)
			os.Exit(1)
		}
		for _, pin := range existingPins {
			if pin.Command == command {
				fmt.Fprintln(os.Stderr, "Error: A command with the same text already exists.")
				os.Exit(1)
			}
		}

		err = database.AddPin(db, command)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error adding pin:", err)
			os.Exit(1)
		}

		fmt.Println("Command pinned successfully!")
	},
}

func getCommand(stdin io.Reader, args []string) string {
	reader := bufio.NewReader(stdin)
	command, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	return strings.TrimSpace(command)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
