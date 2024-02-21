package cmd

import (
	"bufio"
	"fmt"
	"github.com/osean-man/pinner/internal/database"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new pinned command",
	Long: `Add a new pinned command to your list of frequently used commands. You can also pipe in a command with echo:
echo "ls -la" | pinner add`,

	Run: func(cmd *cobra.Command, args []string) {
		command := getCommand(os.Stdin, args)

		existingPins, err := database.GetPins(db)
		if err != nil {
			fmt.Printf("Error fetching pins: %s", err)
			os.Exit(1)
		}
		for _, pin := range existingPins {
			if pin.Command == command {
				fmt.Println("Error: A command with the same text already exists.")
				os.Exit(1)
			}
		}

		err = database.AddPin(db, command)
		if err != nil {
			fmt.Printf("Error adding pin: %s", err)
			os.Exit(1)
		}

		fmt.Println("Command pinned successfully!")
	},
}

func getCommand(stdin io.Reader, args []string) string {
	fmt.Print("Enter the command you want to pin: ")
	reader := bufio.NewReader(stdin)
	command, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading input: %s", err)
		os.Exit(1)
	}

	return strings.TrimSpace(command)
}

func init() {
	rootCmd.AddCommand(addCmd)
}
